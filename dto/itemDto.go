package dto

type ItemUpdateDTO struct {
	LineItemID uint64 `json:"lineitemid" form:"lineitemid" binding:"required"`
	ItemCode   string `json:"itemcode" form:"itemcode" binding:"required"`
	ItemName   string `json:"itemname" form:"itemname" binding:"required"`
	Quantity   int    `json:"quantity" form:"quantity" binding:"required"`
	OrderID    uint64 `json:"order_id,omitempty" form:"order_id, omitempty"`
}

type ItemCreateDTO struct {
	ItemCode string `json:"itemcode" form:"itemcode" binding:"required"`
	ItemName string `json:"itemname" form:"itemname" binding:"required"`
	Quantity int    `json:"quantity" form:"quantity" binding:"required"`
	OrderID  uint64 `json:"order_id,omitempty" form:"order_id, omitempty"`
}
