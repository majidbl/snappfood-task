package controllers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"task/service"
	"task/service/dto"
)

// AssignDelay godoc
// @Summary assign delay
// @Description assign delay of an order that delivery is passed to an agent
// @Tags Delay
// @Accept json
// @Produce json
// @Param AssignDelayRequest body AssignDelayRequest true "necessary item for assign delay"
// @Success 200 {object} AssignDelayResponse
// @Failure	400 {object} AssignDelayResponse "some field is invalid"
// @Failure	500 {object} AssignDelayResponse "other error"
// @Router /api/v1/delay/assign [post]
func AssignDelay() echo.HandlerFunc {
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

		res, err := service.AssignDelay(c, request)
		if err != nil {
			log.Warn(err.Error())
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
