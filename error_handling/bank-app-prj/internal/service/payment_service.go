package service

import (
	"awesomeProject/error_handling/bank-app-prj/internal/domain"
	"fmt"
)

type BankService struct {
	BankaAdi string
}

func (b *BankService) Process(amount float64) error {

	fmt.Println("Banka adi : ", b.BankaAdi)
	if amount > 1000 {
		err := domain.ErrBankSystemDown
		return fmt.Errorf("bank gateway is not responding : %w", err)
	}
	if amount > 500 {
		return &domain.InsufficientFundsError{
			Balance:   500,
			Requested: amount,
		}

	}
	return nil
}
