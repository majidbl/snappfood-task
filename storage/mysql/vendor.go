package mysql

import (
	"context"

	"task/models"
)

// GetVendorsTotalDelay to fetch the total delay minutes for each order
func (s Store) GetVendorsTotalDelay(ctx context.Context) ([]models.VendorDelay, error) {
	var result []models.VendorDelay

	// Construct the query using GORM
	res := s.db.Table("vendors").
		Select("vendors.id AS vendor_id, vendors.name AS vendor_name, orders.id AS order_id, IFNULL(SUM(delay_reports.delivery_time), 0) AS total_delay_minutes").
		Where("WHERE date >= curdate() - INTERVAL DAYOFWEEK(curdate())+6 DAY AND date < curdate() - INTERVAL DAYOFWEEK(curdate())-1 DAY").
		Joins("LEFT JOIN delay_reports ON orders.id = delay_reports.vendor_id").
		Group("vendors.id").
		Scan(&result)

	return result, res.Error
}
