package service

import (
	"go-grpc-demo/internal/model"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (OrderService) CreateOrder(userID, productID int64, amount int32) (model.Order, error) {
	return model.Order{}, nil
}
