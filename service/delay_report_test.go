package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"task/models"
	"task/storage/mysqlstore/mocks"
)

func TestAppointmentService_GetCountAppointments(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var ctx context.Context
	mockVendorRepo := mocks.NewMockIVendor(mockCtrl)
	mockVendorRepo.EXPECT().GetVendorsTotalDelay(ctx).Return([]models.VendorDelay{
		{
			VendorID:          1,
			VendorName:        "vendor-1",
			OrderId:           1,
			TotalDelayMinutes: 15,
		},
		{
			VendorID:          2,
			VendorName:        "vendor-2",
			OrderId:           2,
			TotalDelayMinutes: 30,
		},
	},
		nil)

	delayService := service{vendorDb: mockVendorRepo}
	response, _ := delayService.ReportDelay(ctx)
	if len(response) != 2 {
		t.Errorf("error reporting")
	}
}
