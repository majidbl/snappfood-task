package service

import (
	"context"
	"fmt"
	"log"
	"task/storage/mysqlstore"
	"time"

	"task/dto"
	"task/models"
	"task/storage"
	"task/storage/queue"
	"task/util"
)

func (s service) DelayReport(ctx context.Context, request dto.DelayReportRequest) (dto.DelayReportResponse, error) {
	//store := mysql.NewStore()

	var res dto.DelayReportResponse

	err := s.db.Transaction(
		ctx,
		func(c context.Context, store mysqlstore.IStore) error {
			// Check order exists and it's delivery time passed
			order, err := store.Order().GetOrderById(c, request.OrderId)
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

			// check if order not delayed return error
			if !order.CreatedAt.Add(time.Minute * time.Duration(order.DeliveryTime)).Before(time.Now()) {
				res.Code = models.ErrCode[models.OrderNotDelayedError]
				res.Message = models.OrderNotDelayedError
				return util.NewError(res.Code, res.Message)
			}

			// Check order delay report exists, and if exists its status is reviewed
			var orderDelayReportErr error

			res, orderDelayReportErr = CheckOrderDelayReport(c, store, order.ID)
			if orderDelayReportErr != nil {
				return orderDelayReportErr
			}

			// Check order trips exists and it's status is ok
			orderTrips, getOrderTripsErr := store.Trip().GetOrderTrip(c, request.OrderId)
			if getOrderTripsErr != nil && getOrderTripsErr.Error() != models.OrderTripsNotFountError {
				res.Code = models.ErrCode[models.OrderTripsNotFountError]
				res.Message = models.OrderTripsNotFountError
				return util.NewError(res.Code, res.Message)
			}

			// if the order has no active trips, then immediately we need to create a delay report
			if getOrderTripsErr != nil && getOrderTripsErr.Error() == storage.NotFound {
				queueErr := queue.OrderQueueManger.Enqueue(c, order)
				if queueErr != nil {
					return queueErr
				}

				return store.DelayReport().CreateDelayReport(
					c,
					&models.DelayReport{
						OrderID:  request.OrderId,
						VendorID: order.VendorId,
						Status:   models.ReportRegistered,
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

				return store.DelayReport().CreateDelayReport(
					ctx,
					&models.DelayReport{
						OrderID:      order.ID,
						VendorID:     order.VendorId,
						DeliveryTime: delayTime,
						Status:       models.ReportRegistered,
					})
			}

			res.Code = models.ErrCode[models.OrderDeliveredError]
			res.Message = models.OrderDeliveredError
			return util.NewError(res.Code, res.Message)
		})

	return res, err
}

func CheckOrderDelayReport(c context.Context, store mysqlstore.IStore, orderId uint) (dto.DelayReportResponse, error) {
	var res dto.DelayReportResponse
	orderDelayReport, orderDelayReportErr := store.DelayReport().GetOrderDelayReport(c, orderId)
	if orderDelayReportErr != nil && orderDelayReportErr.Error() != storage.NotFound {
		log.Println(orderDelayReportErr.Error())
		res.Code = models.ErrCode[models.InternalErrorError]
		res.Message = models.InternalErrorError
		return res, util.NewError(res.Code, res.Message)
	}

	validReportDelayStatus := []string{models.ReportRegistered, models.ReportAssigned}
	if util.In(orderDelayReport.Status, validReportDelayStatus) {
		res.Code = models.ErrCode[models.OpenOrderDelayReportProcessError]
		res.Message = models.OpenOrderDelayReportProcessError
		return res, util.NewError(res.Code, res.Message)
	}

	return res, nil
}
