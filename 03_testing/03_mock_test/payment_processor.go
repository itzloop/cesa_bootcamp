package main

// PaymentProcessor defines the interface for processing payments
type PaymentProcessor interface {
	ProcessPayment(amount float64, accountID string) (string, error)
}
