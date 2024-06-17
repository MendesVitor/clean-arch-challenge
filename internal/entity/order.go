package entity

import "time"

type Order struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}
