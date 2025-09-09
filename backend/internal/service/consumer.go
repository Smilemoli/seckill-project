// file: backend/internal/service/consumer.go
package service

import (
	"encoding/json"
	"log"
	"seckill-project/backend/config"
	"seckill-project/backend/internal/dao"
	"seckill-project/backend/internal/model"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

// handleOrder 是真正处理订单创建的函数
func handleOrder(msg OrderMessage) {
	log.Printf("Received an order message: %+v", msg)

	// 查询活动信息，主要是为了获取商品ID和价格
	activity, err := dao.GetActivityByID(msg.ActivityID)
	if err != nil {
		log.Printf("Error getting activity %d: %v", msg.ActivityID, err)
		return
	}

	// 开启数据库事务
	err = dao.DB.Transaction(func(tx *gorm.DB) error {
		// a. 扣减数据库中的库存 (为了最终一致性)
		if err := dao.DecreaseActivityStock(tx, msg.ActivityID, 1); err != nil {
			// 注意：这里可能因为消费者处理速度慢于Redis而导致库存不足的错误
			// 在真实系统中，需要有更完善的补偿和监控机制
			return err
		}

		// b. 创建订单
		order := &model.Order{
			UserID:     msg.UserID,
			ActivityID: msg.ActivityID,
			ProductID:  activity.ProductID,
			TotalPrice: activity.SeckillPrice,
			Status:     2, // 1-处理中, 2-成功, 3-失败
		}
		if err := dao.CreateOrder(tx, order); err != nil {
			return err
		}

		log.Printf("Order created successfully for user %d and activity %d", msg.UserID, msg.ActivityID)
		return nil
	})

	if err != nil {
		log.Printf("Failed to process order in transaction: %v", err)
		// 可以在这里加入重试或失败记录逻辑
	}
}

// StartOrderConsumer 启动订单消费者
func StartOrderConsumer() {
	conn, err := amqp.Dial(config.Conf.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		dao.SeckillOrderQueue, // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 设置 prefetch count 为 1，确保一个 worker 一次只处理一条消息
	// 这在有多个消费者副本时可以做到简单的负载均衡
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack: 改为 false，我们需要手动确认
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 使用一个 channel 来永远阻塞，防止 main goroutine 退出
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var msg OrderMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error unmarshalling message: %s", err)
				// 消息格式错误，直接确认掉，防止队列阻塞
				d.Ack(false)
				continue
			}

			// 处理订单
			handleOrder(msg)

			// 手动发送确认回执
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
