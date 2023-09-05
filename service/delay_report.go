package service

import (
	"context"

	"task/models"
)

func (s service) ReportDelay(ctx context.Context) ([]models.VendorDelay, error) {
	return s.vendorDb.GetVendorsTotalDelay(ctx)
}
