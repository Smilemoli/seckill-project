package service

import (
	"seckill-project/backend/internal/dao"
	"seckill-project/backend/internal/model"
)

// CreateProduct 创建商品
func CreateProduct(p *model.Product) error {
	return dao.CreateProduct(p)
}
