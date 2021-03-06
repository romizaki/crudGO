package dto

type OrderUpdateDTO struct {
	OrderID      uint64 `json:"lineitemid" form:"lineitemid" binding:"required"`
	CustomerName string `json:"customername" form:"customername" binding:"required"`
	OrderAt      string `json:"orderat" form:"orderat" binding:"required"`
	Items        int    `json:"items" form:"items" binding:"required"`
}

type OrderCreateDTO struct {
	OrderID      uint64 `json:"lineitemid" form:"lineitemid" binding:"required"`
	CustomerName string `json:"customername" form:"customername" binding:"required"`
	OrderAt      string `json:"orderat" form:"orderat" binding:"required"`
	Items        int    `json:"items" form:"items" binding:"required"`
}
