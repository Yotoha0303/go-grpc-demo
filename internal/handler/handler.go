package handler

import (
	"go-grpc-demo/internal/model"
	"go-grpc-demo/internal/response"
	"go-grpc-demo/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

type OrderService interface {
	CreateOrder(userID, productID int64, amount int64) (model.Order, error)
	GetOrderByID(orderID int64) (*model.Order, error)
}

var _ OrderService = (*service.OrderService)(nil)

func (o *OrderHandler) CreateOrderHandler(c *gin.Context) {
	// var req request.CreateOrderRequest
	// if err := c.ShouldBindJSON(req); err != nil {
	// 	response.Fail(c, http.StatusBadRequest, err.Error())
	// }
	// order, err := o.orderService.CreateOrder(req.IdempotencyKey, req.Items[0].ProductID, req.Items[0].Quantity)
	// if err != nil {
	// 	response.Fail(c, http.StatusInternalServerError, err.Error())
	// }
	// response.Success(c, order)
}

func parsePositiveID(c *gin.Context, paramName string) (int64, bool) {

	id, err := strconv.ParseInt(c.Param(paramName), 10, 64)
	if err != nil || id <= 0 {
		response.Fail(c, http.StatusBadRequest, err.Error())
		return 0, false
	}
	return id, true
}

func (o *OrderHandler) GetOrderHandler(c *gin.Context) {
	id, ok := parsePositiveID(c, "id")
	if !ok {
		return
	}

	order, err := o.orderService.GetOrderByID(id)
	if err != nil {
		response.Fail(c, http.StatusNotFound, err.Error())
	}

	response.Success(c, order)
}
