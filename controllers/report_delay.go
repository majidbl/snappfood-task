package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"task/models"
	"task/storage"
	"task/storage/mysql"
	"task/storage/queue"
	"task/util"
)

type DelayReportRequest struct {
	OrderId uint `json:"orderId" validate:"required"`
}

type DelayReportResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReportDelay godoc
// @Summary report delay
// @Description report delay of a order that delivery is passed
// @Tags Delay
// @Accept json
// @Produce json
// @Param DelayReportRequest body DelayReportRequest true "necessary item for reporting delay"
// @Success 200 {object} DelayReportResponse
// @Failure	400 {object} DelayReportResponse "some field is invalid"
// @Failure	500 {object} DelayReportResponse "other error"
// @Router /api/v1/delay/report [post]
func ReportDelay() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request DelayReportRequest

		// add required validator tag
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		c := ctx.Request().Context()
		store := mysql.NewStore()

		var res DelayReportResponse

		err := store.Transaction(
			c,
			func(ctx context.Context, store mysql.Store) error {
				// Check order exists and it's delivery time passed
				order, err := store.GetDelayedOrder(c, request.OrderId)
				if err != nil {
					if err.Error() == storage.NotFound {
						res.Code = models.ErrCode[models.OrderNotFountError]
						res.Message = models.OrderNotFountError
						return util.NewError(res.Code, res.Message)
					}

					log.Println(err.Error())
					res.Code = models.ErrCode[models.InternalErrorError]
					res.Message = models.InternalErrorError
					return util.NewError(res.Code, res.Message)

				}

				if !order.CreatedAt.Add(time.Minute * time.Duration(order.DeliveryTime)).Before(time.Now()) {
					res.Code = models.ErrCode[models.OrderNotDelayedError]
					res.Message = models.OrderNotDelayedError
					return util.NewError(res.Code, res.Message)
				}

				// Check order delay report exists and it's status is ok
				orderDelayReport, orderDelayReportErr := store.GetOrderDelayReport(c, request.OrderId)
				if orderDelayReportErr != nil && orderDelayReportErr.Error() != storage.NotFound {
					log.Println(orderDelayReportErr.Error())
					res.Code = models.ErrCode[models.InternalErrorError]
					res.Message = models.InternalErrorError
					return util.NewError(res.Code, res.Message)
				}

				validReportDelayStatus := []string{models.ReportRegistered, models.ReportAssigned}
				if util.In(orderDelayReport.Status, validReportDelayStatus) {
					res.Code = models.ErrCode[models.OpenOrderDelayReportProcessError]
					res.Message = models.OpenOrderDelayReportProcessError
					return util.NewError(res.Code, res.Message)
				}

				// Check order trips exists and it's status is ok
				orderTrips, getOrderTripsErr := store.GetOrderTrip(c, request.OrderId)
				if getOrderTripsErr != nil && getOrderTripsErr.Error() != models.OrderTripsNotFountError {
					res.Code = models.ErrCode[models.OrderTripsNotFountError]
					res.Message = models.OrderTripsNotFountError
					return util.NewError(res.Code, res.Message)
				}

				// if the order has no active trips, then immediately we need to create a delay report
				if getOrderTripsErr != nil && getOrderTripsErr.Error() == storage.NotFound {
					return store.CreateDelayReport(
						c,
						&models.DelayReport{
							OrderID: request.OrderId,
							Status:  models.ReportRegistered,
						})
				}

				// check trips status for make decision
				validTripStatus := []string{models.ASSIGNED, models.AtVENDOR, models.PICKED}
				if util.In(orderTrips.Status, validTripStatus) {
					delayTime := util.MockDelay()
					res.Code = models.ErrCode[models.OperationSuccess]
					res.Message = fmt.Sprintf(
						"your order deliver at %s",
						time.Now().Add(time.Minute*time.Duration(delayTime)),
					)

					queueErr := queue.OrderQueueManger.Enqueue(c, order)
					if queueErr != nil {
						return queueErr
					}

					return store.CreateDelayReport(
						ctx,
						&models.DelayReport{
							OrderID:      order.ID,
							DeliveryTime: delayTime,
							Status:       models.ReportRegistered,
						})
				}

				return nil
			})

		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
