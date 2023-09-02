package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"task/models"
	"task/storage/mysql"
)

type ReportDelayResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// DelayReport godoc
// @Summary report delay
// @Description report delay of an order that delivery is passed
// @Tags Delay
// @Accept json
// @Produce json
// @Success 200 {object} ReportDelayResponse
// @Failure	400 {object} ReportDelayResponse "some field is invalid"
// @Failure	500 {object} ReportDelayResponse "other error"
// @Router /api/v1/delay/report [get]
func DelayReport() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		c := ctx.Request().Context()

		var response ReportDelayResponse
		res, err := mysql.NewStore().GetVendorsTotalDelay(c)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		response.Result = res
		response.Code = models.ErrCode[models.OperationSuccess]
		response.Message = models.OperationSuccess
		return ctx.JSON(http.StatusOK, response)
	}
}
