package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"task/models"
	"task/service/mocks"
)

func TestDelayDelayReport(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := c.Request().Context()

	notificationsService := mocks.NewMockINotificationService(mockCtrl)
	notificationsService.EXPECT().ReportDelay(ctx).Return([]models.VendorDelay{
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

	controller := Controller{notificationsService}
	err := controller.DelayReport()(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
