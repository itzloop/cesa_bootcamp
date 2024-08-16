package main

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentServiceChargeAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProcessor := NewMockPaymentProcessor(ctrl)
	paymentService := NewPaymentService(mockProcessor)
	mockProcessor.EXPECT().ProcessPayment(100.0, "account123").Return("tx123", nil)
	transactionID, err := paymentService.ChargeAccount(100.0, "account123")
	assert.NoError(t, err)
	assert.Equal(t, "tx123", transactionID)

	mockProcessor.EXPECT().ProcessPayment(200.0, "account1234").Return("132123", errors.New("no tx"))
	transactionID, err = paymentService.ChargeAccount(200.0, "account1234")
	assert.Error(t, err)
	fmt.Println(err)
	assert.Equal(t, "", transactionID)
}
