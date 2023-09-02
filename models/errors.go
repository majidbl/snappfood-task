package models

const (
	OrderNotFountError               = "order not found"
	OrderNotDelayedError             = "order not delayed"
	OrderDeliveredError              = "order delivered"
	DelayedOrderNotFoundError        = "delayed order not found"
	OrderTripsNotFountError          = "order trips not found"
	OrderDelayReportNotFountError    = "order delay report not found"
	OpenOrderDelayReportProcessError = "open order delay report process error"
	AgentBusyErr                     = "agent is busy and cannot talk new orders"
	OperationSuccess                 = "operation success"
	InternalErrorError               = "internal error"
)

var ErrCode = map[string]int{
	OrderNotFountError:               100404,
	OrderNotDelayedError:             100400,
	OrderDeliveredError:              100420,
	DelayedOrderNotFoundError:        100410,
	OrderTripsNotFountError:          200404,
	OrderDelayReportNotFountError:    300404,
	OpenOrderDelayReportProcessError: 400420,
	AgentBusyErr:                     500503,
	InternalErrorError:               500,
	OperationSuccess:                 200,
}
