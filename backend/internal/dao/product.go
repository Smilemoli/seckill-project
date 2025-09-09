package dao

import "seckill-project/backend/internal/model"

// CreateProduct 在数据库中创建一个新商品
func CreateProduct(product *model.Product) error {
	return DB.Create(product).Error
}
