package dao

import (
	"context"
	"go-grpc-demo/internal/model"

	"gorm.io/gorm"
)

func GetOrderByID(ctx context.Context, db *gorm.DB, id int64) (*model.Order, error) {
	var order model.Order
	return &order, db.WithContext(ctx).Where("id = ?", id).First(&order).Error
}

func CreateOrder(ctx context.Context, db *gorm.DB, order *model.Order) error {
	return db.WithContext(ctx).Create(&order).Error
}
