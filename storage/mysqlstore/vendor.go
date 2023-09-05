package mysqlstore

import (
	"context"

	"gorm.io/gorm"

	"task/models"
)

type IVendor interface {
	GetVendorsTotalDelay(ctx context.Context) ([]models.VendorDelay, error)
}

type vendor struct {
	db *gorm.DB
}

func NewVendor(db *gorm.DB) IVendor {
	return &vendor{db: db}
}

// GetVendorsTotalDelay to fetch the total delay minutes for each order
func (v vendor) GetVendorsTotalDelay(ctx context.Context) ([]models.VendorDelay, error) {
	var result []models.VendorDelay

	// Construct the query using GORM
	res := v.db.Table("vendors").
		Select("vendors.id AS vendor_id, vendors.name AS vendor_name, IFNULL(SUM(delay_reports.delivery_time), 0) AS total_delay_minutes").
		Joins("LEFT JOIN delay_reports ON vendors.id = delay_reports.vendor_id").
		Where("delay_reports.created_at >= curdate() - INTERVAL DAYOFWEEK(curdate())+6 DAY AND delay_reports.created_at < curdate() - INTERVAL DAYOFWEEK(curdate())-1 DAY").
		Group("vendors.id").
		Scan(&result)

	return result, res.Error
}
