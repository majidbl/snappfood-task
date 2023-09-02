package service

import (
	"context"

	"task/models"
	"task/storage/mysql"
)

func ReportDelay(ctx context.Context) ([]models.VendorDelay, error) {
	return mysql.NewStore().GetVendorsTotalDelay(ctx)
}
