package service

import (
	"context"

	"task/dto"
	"task/models"
	"task/storage/mysqlstore"
)

type INotificationService interface {
	AssignDelay(c context.Context, request dto.AssignDelayRequest) (dto.AssignDelayResponse, error)
	ReportDelay(ctx context.Context) ([]models.VendorDelay, error)
	DelayReport(ctx context.Context, request dto.DelayReportRequest) (dto.DelayReportResponse, error)
}

type service struct {
	db       mysqlstore.ITransaction
	vendorDb mysqlstore.IVendor
}

func NewService(db mysqlstore.ITransaction, vendorDb mysqlstore.IVendor) INotificationService {
	return &service{
		db:       db,
		vendorDb: vendorDb,
	}
}
