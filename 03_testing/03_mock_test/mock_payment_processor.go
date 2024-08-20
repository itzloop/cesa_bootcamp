// Code generated by MockGen. DO NOT EDIT.
// Source: payment_processor.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPaymentProcessor is a mock of PaymentProcessor interface.
type MockPaymentProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentProcessorMockRecorder
}

// MockPaymentProcessorMockRecorder is the mock recorder for MockPaymentProcessor.
type MockPaymentProcessorMockRecorder struct {
	mock *MockPaymentProcessor
}

// NewMockPaymentProcessor creates a new mock instance.
func NewMockPaymentProcessor(ctrl *gomock.Controller) *MockPaymentProcessor {
	mock := &MockPaymentProcessor{ctrl: ctrl}
	mock.recorder = &MockPaymentProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentProcessor) EXPECT() *MockPaymentProcessorMockRecorder {
	return m.recorder
}

// ProcessPayment mocks base method.
func (m *MockPaymentProcessor) ProcessPayment(amount float64, accountID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPayment", amount, accountID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessPayment indicates an expected call of ProcessPayment.
func (mr *MockPaymentProcessorMockRecorder) ProcessPayment(amount, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPayment", reflect.TypeOf((*MockPaymentProcessor)(nil).ProcessPayment), amount, accountID)
}