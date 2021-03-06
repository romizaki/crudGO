package repository

import (
	"github.com/romizaki/crudGO/entity"
	"gorm.io/gorm"
)

// ItemRepository ...
type ItemRepository interface {
	InsertItem(b entity.Item) entity.Item
}

type itemConnection struct {
	connection *gorm.DB
}

func (db *itemConnection) InsertItem(b entity.Item) entity.Item {
	db.connection.Save(&b)
	db.connection.Preload("Order").Find(&b)
	return b
}
