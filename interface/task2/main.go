package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrInvalidAmount       = errors.New("некорректная сумма платежа")
	ErrProviderUnavailable = errors.New("провайдер недоступен")
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type Sberbank struct {
	APIKey string
}

type Tbank struct {
	APIKey string
}

type Alfabank struct {
	APIKey string
}

func NewSberbank(apiKey string) *Sberbank {
	return &Sberbank{APIKey: apiKey}
}

func (s *Sberbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Intn(100) >= 80 {
		return ErrProviderUnavailable
	}

	return nil
}

func NewTbank(apiKey string) *Tbank {
	return &Tbank{APIKey: apiKey}
}

func (t *Tbank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Intn(100) >= 90 {
		return ErrProviderUnavailable
	}

	return nil
}

func NewAlfabank(apiKey string) *Alfabank {
	return &Alfabank{APIKey: apiKey}
}

func (a *Alfabank) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if rand.Intn(100) >= 75 {
		return ErrProviderUnavailable
	}

	return nil
}

func main() {
	sberbank := PaymentProcessor(NewSberbank("qwerty123"))
	tbank := PaymentProcessor(NewTbank("asdfgh456"))
	alfabank := PaymentProcessor(NewAlfabank("zxcvbn789"))

	fmt.Println("Sberbank: ", sberbank.ProcessPayment(10))
	fmt.Println("Tbank: ", tbank.ProcessPayment(-5))
	fmt.Println("Alfabank: ", alfabank.ProcessPayment(15))
}
