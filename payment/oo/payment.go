package oo

type PaymentOption interface {
	ProcessPayment(float32) bool
}
