package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not null" json:"customerName"`
	Items        []Item    `json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
}
