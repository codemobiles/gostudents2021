package model

import "time"

// https://gorm.io/docs/query.html
type Transaction struct {
	ID            uint    `gorm:"primary_key"`
	Total         float64 `json:"total"`
	Paid          float64 `json:"paid"`
	Change        float64 `json:"change"`
	PaymentType   string  `json:"payment_type"`
	PaymentDetail string  `json:"payment_detail"`
	OrderList     string  `json:"order_list"`
	StaffID       string  `json:"staff_id"`
	CreatedAt     time.Time
}
