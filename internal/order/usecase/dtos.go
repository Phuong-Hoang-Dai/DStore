package usecase

import "github.com/Phuong-Hoang-Dai/DStore/internal/order"

type OrderDTO struct {
	ProductId int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderId   int     `json:"orderId"`
}

type OrderResponeDTO struct {
	Price float64 `json:"price"`
}

func MapOrderResponeDTOtoOrderDTO(oR OrderResponeDTO, o *OrderDTO) {
	o.Price = oR.Price
}

func MapOrderDTOtoOrderItem(oD OrderDTO, o *order.OrderItem) {
	o.ProductId = oD.ProductId
	o.Quantity = oD.Quantity
	o.Price = oD.Price
	o.OrderId = oD.OrderId
}

func MapOrderItemtoOrderDTO(o order.OrderItem, oD *OrderDTO) {
	oD.ProductId = o.ProductId
	oD.Quantity = o.Quantity
	oD.Price = o.Price
	oD.OrderId = o.OrderId
}
