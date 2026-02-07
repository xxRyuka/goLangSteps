package domain

import (
	"errors"
	"fmt"
)

// 1. Sentinel Error (Sabit Hata)
// Başlangıç değerini atamazsan nil olur!
var ErrBankSystemDown = errors.New("banka sistemi yanit vermiyor")

type InsufficientFundsError struct {
	Balance   float64
	Requested float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("yetersiz bakiye! Mevcut: %.2f, İstenen: %.2f", e.Balance, e.Requested)
}
