// Code generated by MockGen. DO NOT EDIT.
// Source: delay_report.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "task/models"

	gomock "github.com/golang/mock/gomock"
)

// MockIDelayReport is a mock of IDelayReport interface.
type MockIDelayReport struct {
	ctrl     *gomock.Controller
	recorder *MockIDelayReportMockRecorder
}

// MockIDelayReportMockRecorder is the mock recorder for MockIDelayReport.
type MockIDelayReportMockRecorder struct {
	mock *MockIDelayReport
}

// NewMockIDelayReport creates a new mock instance.
func NewMockIDelayReport(ctrl *gomock.Controller) *MockIDelayReport {
	mock := &MockIDelayReport{ctrl: ctrl}
	mock.recorder = &MockIDelayReportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDelayReport) EXPECT() *MockIDelayReportMockRecorder {
	return m.recorder
}

// CreateDelayReport mocks base method.
func (m *MockIDelayReport) CreateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDelayReport", ctx, delayReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDelayReport indicates an expected call of CreateDelayReport.
func (mr *MockIDelayReportMockRecorder) CreateDelayReport(ctx, delayReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDelayReport", reflect.TypeOf((*MockIDelayReport)(nil).CreateDelayReport), ctx, delayReport)
}

// GetOrderDelayReport mocks base method.
func (m *MockIDelayReport) GetOrderDelayReport(ctx context.Context, orderId uint) (models.DelayReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDelayReport", ctx, orderId)
	ret0, _ := ret[0].(models.DelayReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDelayReport indicates an expected call of GetOrderDelayReport.
func (mr *MockIDelayReportMockRecorder) GetOrderDelayReport(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDelayReport", reflect.TypeOf((*MockIDelayReport)(nil).GetOrderDelayReport), ctx, orderId)
}

// UpdateDelayReport mocks base method.
func (m *MockIDelayReport) UpdateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDelayReport", ctx, delayReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDelayReport indicates an expected call of UpdateDelayReport.
func (mr *MockIDelayReportMockRecorder) UpdateDelayReport(ctx, delayReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDelayReport", reflect.TypeOf((*MockIDelayReport)(nil).UpdateDelayReport), ctx, delayReport)
}
