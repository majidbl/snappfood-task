package controllers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"task/dto"
	"task/util"
)

// AssignDelay godoc
// @Summary assign delay
// @Description assign delay of an order that delivery is passed to an agent
// @Tags Delay
// @Accept json
// @Produce json
// @Param AssignDelayRequest body dto.AssignDelayRequest true "necessary item for assign delay"
// @Success 200 {object} dto.AssignDelayResponse
// @Failure	400 {object} dto.AssignDelayResponse "some field is invalid"
// @Failure	500 {object} dto.AssignDelayResponse "other error"
// @Router /api/v1/delay/assign [post]
func (ctrl Controller) AssignDelay() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := dto.AssignDelayRequest{}
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		c := ctx.Request().Context()

		res, err := ctrl.service.AssignDelay(c, request)
		if err != nil {
			log.Warn(err.Error())
			return ctx.JSON(http.StatusInternalServerError, util.CastError(err))
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
