package request

type OrderRequest struct {
	OrderID int64 `json:"order_id" binding:"required"`
}

// type CreateOrderRequest struct {
// 	UserID    int64
// 	ProductID int64
// 	Amount    int64
// }

type CreateOrderRequest struct {
	IdempotencyKey string                   `json:"idempotency_key" binding:"required,max=128"`
	Items          []CreateOrderItemRequest `json:"items" binding:"required,min=1,dive"`
}

type CreateOrderItemRequest struct {
	ProductID int64 `json:"product_id" binding:"required,gt=0"`
	Quantity  int64 `json:"quantity" binding:"required,gt=0"`
}
