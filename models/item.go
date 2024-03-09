package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"not null" json:"itemCode"`
	Description string `gorm:"not null" json:"description"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint
}
