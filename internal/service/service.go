package service

import (
	"context"
	"go-grpc-demo/internal/dao"
	"go-grpc-demo/internal/model"

	"gorm.io/gorm"
)

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (o OrderService) CreateOrder(userID, productID int64, amount int64) (model.Order, error) {
	var order model.Order
	return order, dao.CreateOrder(context.Background(), o.db, &order)
}

func (o OrderService) GetOrderByID(orderID int64) (*model.Order, error) {
	return dao.GetOrderByID(context.Background(), o.db, orderID)
}
