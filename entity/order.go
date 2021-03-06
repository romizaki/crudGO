package entity

//Order database
type Order struct {
	OrderID      uint64  `gorm:"primary_key:auto_increment" json:"id"`
	CustomerName string  `gorm:"type:varchar(255)" json:"name"`
	OrderedAt    string  `gorm:"type:varchar(255)" json:"orderAt"`
	Items        *[]Item `json:"items,omitempty"`
}
