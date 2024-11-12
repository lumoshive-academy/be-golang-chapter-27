package model

type Order struct {
	ID             uint        `json:"id"`
	Customer       Customer    `json:"customer"`
	Address        Address     `json:"address"`
	OrderItem      []OrderItem `json:"order_item"`
	Status         string      `json:"status"`
	PaymentMethod  string      `json:"payment_method"`
	TotalAmount    float32     `json:"total_amount"`
	DiscountAmount float32     `json:"dicount_amount"`
	TimeStamp
}

type OrderItem struct {
	ID       uint `json:"id"`
	Book     Book `json:"book"`
	Quantity int  `json:"quantity"`
	TimeStamp
}

type ReviewOrder struct {
	ID     uint    `json:"id"`
	Book   Book    `json:"book"`
	Order  Order   `json:"order"`
	Rating float32 `json:"rating"`
	TimeStamp
}
