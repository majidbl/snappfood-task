package mysql2

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"task/storage/mysql"
)

type Fn func(context.Context, IStore) error

type store struct {
	agent  IAgent
	order  IOrder
	vendor IVendor
}

type storex struct {
	db *gorm.DB
}

type IStore interface {
	Agent() IAgent
	Order() IOrder
	Vendor() IVendor
}

type IStorex interface {
	Transaction(ctx context.Context, fn Fn) (err error)
}

func (s store) Agent() IAgent {
	return s.agent
}

func (s store) Order() IOrder {
	return s.order
}

func (s store) Vendor() IVendor {
	return s.vendor
}

func NewStore() IStorex {
	return &storex{}
}

func (s storex) Transaction(ctx context.Context, fn Fn) (err error) {
	tx := mysql.NewDB().Begin()

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()

	newStore := store{
		agent:  NewAgent(tx),
		order:  NewOrder(tx),
		vendor: NewVendor(tx),
	}

	err = fn(ctx, newStore)
	return err
}
