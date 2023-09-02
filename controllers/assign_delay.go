package controllers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"task/models"
	"task/storage"
	"task/storage/mysql"
	"task/storage/queue"
	"task/util"
)

type AssignDelayRequest struct {
	AgentId int64 `json:"agentId" validate:"required"`
}

type AssignDelayResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	OrderId uint   `json:"orderId"`
}

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
		var request AssignDelayRequest

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

		var res AssignDelayResponse

		err := store.Transaction(
			c,
			func(ctx context.Context, store mysql.Store) error {
				agent, err := store.GetAgent(c, request.AgentId)
				if err != nil {
					return err
				}

				if agent.Status == models.Busy {
					res.Code = models.ErrCode[models.AgentBusyErr]
					res.Message = models.AgentBusyErr
					return util.NewError(res.Code, res.Message)
				}

				order, queueErr := queue.OrderQueueManger.Dequeue(c)
				if queueErr != nil && queueErr.Error() != queue.EmptyQueue {
					log.Warn(queueErr.Error())
					res.Code = models.ErrCode[models.InternalErrorError]
					res.Message = models.InternalErrorError + ":" + queueErr.Error()
					return util.NewError(res.Code, res.Message)
				}

				// if Order Queue Manager Was empty, we can check a database as reference
				if queueErr.Error() == queue.EmptyQueue {
					orders, ordersDelayedReportErr := store.GetDelayedOrders(c)
					if err != nil {
						log.Warn(ordersDelayedReportErr.Error())
						res.Code = models.ErrCode[models.InternalErrorError]
						res.Message = models.InternalErrorError
						return util.NewError(res.Code, res.Message)
					}

					if orders == nil {
						res.Code = models.ErrCode[models.DelayedOrderNotFoundError]
						res.Message = models.DelayedOrderNotFoundError
						return util.NewError(res.Code, res.Message)
					}

					for _, m := range orders {
						enqueueErr := queue.OrderQueueManger.Enqueue(c, m)
						if enqueueErr != nil {
							log.Warn(enqueueErr.Error())
							res.Code = models.ErrCode[models.InternalErrorError]
							res.Message = models.InternalErrorError
							return util.NewError(res.Code, res.Message)
						}
					}
				}

				// Check order delay report exists and it's status is ok
				orderDelayReport, orderDelayReportErr := store.GetOrderDelayReport(c, order.ID)
				if orderDelayReportErr != nil && orderDelayReportErr.Error() != storage.NotFound {
					log.Warn(orderDelayReportErr.Error())
					res.Code = models.ErrCode[models.InternalErrorError]
					res.Message = models.InternalErrorError
					return util.NewError(res.Code, res.Message)
				}

				validReportDelayStatus := []string{models.ReportAssigned}
				if util.In(orderDelayReport.Status, validReportDelayStatus) {
					res.Code = models.ErrCode[models.OpenOrderDelayReportProcessError]
					res.Message = models.OpenOrderDelayReportProcessError
					return util.NewError(res.Code, res.Message)
				}

				agent.Status = models.Busy
				order.State = models.OrderReview
				orderDelayReport.Status = models.ReportAssigned
				orderDelayReport.AgentId = agent.ID

				updateOrderErr := store.UpdateOrder(c, &order)
				updateOrderDelayReportErr := store.UpdateDelayReport(c, &orderDelayReport)
				updateAgentErr := store.UpdateAgent(c, &agent)

				if updateOrderErr != nil || updateAgentErr != nil || updateOrderDelayReportErr != nil {
					enqueueErr := queue.OrderQueueManger.Enqueue(c, order)
					if enqueueErr != nil {
						log.Warn(enqueueErr.Error())
					}

					log.Warn("updateOrderErr: ", updateOrderErr)
					log.Warn("updateAgentErr", updateAgentErr)
					log.Warn("updateOrderDelayReportErr", orderDelayReportErr)
					return util.NewError(models.ErrCode[models.InternalErrorError], models.InternalErrorError)
				}

				res.OrderId = order.ID
				return err
			})

		if err != nil {
			log.Warn(err.Error())
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	}
}
