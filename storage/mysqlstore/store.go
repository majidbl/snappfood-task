package mysqlstore

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Fn func(context.Context, IStore) error

type store struct {
	agent       IAgent
	order       IOrder
	vendor      IVendor
	delayReport IDelayReport
	trip        ITrip
}

type transaction struct {
	db *gorm.DB
}

type IStore interface {
	Agent() IAgent
	Order() IOrder
	Vendor() IVendor
	DelayReport() IDelayReport
	Trip() ITrip
}

type ITransaction interface {
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

func (s store) DelayReport() IDelayReport {
	return s.delayReport
}

func (s store) Trip() ITrip {
	return s.trip
}

func NewStore() ITransaction {
	return &transaction{}
}

func (s transaction) Transaction(ctx context.Context, fn Fn) (err error) {
	tx := DB.Begin()

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr.Error != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()

	newStore := store{
		agent:       NewAgent(tx),
		order:       NewOrder(tx),
		vendor:      NewVendor(tx),
		delayReport: NewDelayReport(tx),
		trip:        NewTrip(tx),
	}

	err = fn(ctx, newStore)
	return err
}
