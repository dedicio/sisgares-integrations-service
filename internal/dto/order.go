package dto

type OrderItemDTO struct {
	ID        string  `json:"id"`
	ProductId string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderDTO struct {
	ID            string         `json:"id"`
	CompanyID     string         `json:"company_id"`
	Items         []OrderItemDTO `json:"items"`
	Discount      float64        `json:"discount"`
	PaymentMethod string         `json:"payment_method"`
}

type OrderMessagingItemDTO struct {
	ID        string  `json:"ID"`
	ProductId string  `json:"ProductID"`
	Quantity  int     `json:"Quantity"`
	Price     float64 `json:"Price"`
}

type OrderMessagingDTO struct {
	ID            string                  `json:"ID"`
	CompanyID     string                  `json:"CompanyId"`
	Items         []OrderMessagingItemDTO `json:"Items"`
	Discount      float64                 `json:"Discount"`
	Status        string                  `json:"Status"`
	PaymentMethod string                  `json:"PaymentMethod"`
}
