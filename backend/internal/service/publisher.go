// 专门负责发消息
package service

import (
	"encoding/json"
	"log"
	"seckill-project/backend/config"
	"seckill-project/backend/internal/dao"

	"github.com/streadway/amqp"
)

type OrderMessage struct {
	UserID     uint `json:"user_id"`
	ActivityID uint `json:"activity_id"`
}

// PublishOrderMessage 发送订单消息到 RabbitMQ
func PublishOrderMessage(msg OrderMessage) {
	conn, err := amqp.Dial(config.Conf.RabbitMQ.URL)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		dao.SeckillOrderQueue, // name
		true,                  // durable: 队列持久化
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息持久化
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
	}
}
