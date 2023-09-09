package controllers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"task/dto"
	"task/util"
)

// ReportDelay godoc
// @Summary report delay
// @Description report delay of a order that delivery is passed
// @Tags Delay
// @Accept json
// @Produce json
// @Param DelayReportRequest body dto.DelayReportRequest true "necessary item for reporting delay"
// @Success 200 {object} dto.DelayReportResponse
// @Failure	400 {object} dto.DelayReportResponse "some field is invalid"
// @Failure	500 {object} dto.DelayReportResponse "other error"
// @Router /api/v1/delay/report [post]
func (ctrl Controller) ReportDelay() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request dto.DelayReportRequest

		// add required validator tag
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		c := ctx.Request().Context()

		res, err := ctrl.service.DelayReport(c, request)
		if err != nil {
			log.Warn(err.Error())
			return ctx.JSON(http.StatusInternalServerError, util.CastError(err))
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
