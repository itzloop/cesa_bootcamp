package main

import (
	"errors"
)

// PaymentService is responsible for handling payments
type PaymentService struct {
	processor PaymentProcessor
}

// NewPaymentService creates a new PaymentService with the given PaymentProcessor
func NewPaymentService(processor PaymentProcessor) *PaymentService {
	return &PaymentService{processor: processor}
}

// ChargeAccount processes a payment for the given account ID
func (s *PaymentService) ChargeAccount(amount float64, accountID string) (string, error) {
	if amount <= 0 {
		return "", errors.New("invalid amount")
	}

	// Use the PaymentProcessor to process the payment
	transactionID, err := s.processor.ProcessPayment(amount, accountID)
	if err != nil {
		return "", err
	}

	return transactionID, nil
}
