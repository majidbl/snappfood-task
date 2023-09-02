package dto

type DelayReportRequest struct {
	OrderId uint `json:"orderId" validate:"required"`
}

type DelayReportResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
