package model

import "time"

type Product struct {
	ID        uint      `json:"id"; gorm:"primary_key"`
	Name      string    `json:"name"`
	Stock     int64     `json:"stock"`
	Price     float64   `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"time"`
}
