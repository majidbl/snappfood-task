package dto

type AssignDelayRequest struct {
	AgentId int64 `json:"agentId" validate:"required"`
}

type AssignDelayResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	OrderId uint   `json:"orderId"`
}
