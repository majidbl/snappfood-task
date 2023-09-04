package mysqlstore

import (
	"context"

	"gorm.io/gorm"

	"task/models"
)

type ITrip interface {
	GetOrderTrip(ctx context.Context, orderId uint) (models.Trip, error)
}

type trip struct {
	db *gorm.DB
}

func NewTrip(db *gorm.DB) ITrip {
	return &trip{db: db}
}

func (t trip) GetOrderTrip(ctx context.Context, orderId uint) (models.Trip, error) {
	var tr models.Trip
	result := t.db.First(&tr, "order_id=?", orderId)
	return tr, result.Error
}
