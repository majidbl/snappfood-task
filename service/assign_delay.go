package service

import (
	"context"
	"task/dto"

	"github.com/labstack/gommon/log"

	"task/models"
	"task/storage"
	"task/storage/mysqlstore"
	"task/storage/queue"
	"task/util"
)

func (s service) AssignDelay(c context.Context, request dto.AssignDelayRequest) (dto.AssignDelayResponse, error) {
	var res dto.AssignDelayResponse

	err := s.db.Transaction(
		c,
		func(ctx context.Context, store mysqlstore.IStore) error {
			agent, err := store.Agent().GetAgent(ctx, request.AgentId)
			if err != nil {
				res.Code = models.ErrCode[models.InternalErrorError]
				res.Message = models.InternalErrorError
				return util.NewError(res.Code, res.Message)
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

			if queueErr.Error() == queue.EmptyQueue {
				res, err = FillQueueWithDelayedOrder(ctx, store)
			}

			// Check order delay report exists and it's status is ok
			orderDelayReport, orderDelayReportErr := store.DelayReport().GetOrderDelayReport(c, order.ID)
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

			// TODO: move to function assign order to agent
			agent.Status = models.Busy
			order.State = models.OrderReview
			orderDelayReport.Status = models.ReportAssigned
			orderDelayReport.AgentId = agent.ID

			updateOrderErr := store.Order().UpdateOrder(c, &order)
			updateOrderDelayReportErr := store.DelayReport().UpdateDelayReport(c, &orderDelayReport)
			updateAgentErr := store.Agent().UpdateAgent(c, &agent)

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

	return res, err
}

func FillQueueWithDelayedOrder(ctx context.Context, store mysqlstore.IStore) (dto.AssignDelayResponse, error) {
	var res dto.AssignDelayResponse
	// if Order Queue Manager Was empty, we can check a database as reference
	orders, ordersDelayedReportErr := store.Order().GetDelayedOrders(ctx)
	if ordersDelayedReportErr != nil {
		log.Warn(ordersDelayedReportErr.Error())
		res.Code = models.ErrCode[models.InternalErrorError]
		res.Message = models.InternalErrorError
		return res, util.NewError(res.Code, res.Message)
	}

	if orders == nil {
		res.Code = models.ErrCode[models.DelayedOrderNotFoundError]
		res.Message = models.DelayedOrderNotFoundError
		return res, util.NewError(res.Code, res.Message)
	}

	for _, m := range orders {
		enqueueErr := queue.OrderQueueManger.Enqueue(ctx, m)
		if enqueueErr != nil {
			log.Warn(enqueueErr.Error())
			res.Code = models.ErrCode[models.InternalErrorError]
			res.Message = models.InternalErrorError
			return res, util.NewError(res.Code, res.Message)
		}
	}

	return res, nil
}
