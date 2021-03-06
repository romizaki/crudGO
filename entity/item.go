package entity

//Item database
type Item struct {
	LineItemID uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ItemCode   string `gorm:"type:varchar(255)" json:"itemCode"`
	ItemName   string `gorm:"type:text" json:"itemName"`
	Quantity   int    `gorm:"type:integer" json:"quantity"`
	OrderID    uint64 `gorm:"not null" json:"orderId"`
	Order      Order  `gorm:"foreignkey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"order"`
}
