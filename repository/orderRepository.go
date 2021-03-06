package repository

import (
	"github.com/romizaki/crudGO/entity"
	"gorm.io/gorm"
)

// OrderRepository ...
type OrderRepository interface {
	InsertOrder(b entity.Order) entity.Order
	UpdateOrder(b entity.Order) entity.Order
	DeleteOrder(b entity.Order)
	AllOrder() []entity.Order
	FindOrderByID(orderID uint64) entity.Order
}

type orderConnection struct {
	connection *gorm.DB
}

func (db *orderConnection) InsertOrder(b entity.Order) entity.Order {
	db.connection.Save(&b)
	db.connection.Preload("Item").Find(&b)
	return b
}

func (db *orderConnection) UpdateOrder(b entity.Order) entity.Order {
	db.connection.Save(&b)
	db.connection.Preload("Item").Find(&b)
	return b
}

func (db *orderConnection) DeleteOrder(b entity.Order) {
	db.connection.Delete(&b)
}

func (db *orderConnection) FindOrderByID(orderID uint64) entity.Order {
	var order entity.Order
	db.connection.Preload("Item").Find(&order, orderID)
	return order
}

func (db *orderConnection) AllOrder() []entity.Order {
	var orders []entity.Order
	db.connection.Preload("Item").Find(&orders)
	return orders
}
