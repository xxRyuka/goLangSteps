package domain

import (
	"fmt"
	"sync"
)

type Wallet struct {
	Balanece float64
	mtx      sync.Mutex
	//Assets   map[string]float64
}

func NewWallet(initialBalance float64) *Wallet {
	return &Wallet{
		Balanece: initialBalance,
	}
}

func (w *Wallet) Deposit(amount float64) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	w.Balanece = w.Balanece + amount
}

func (w *Wallet) ProcessBuy(amount float64) (bool, float64) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.Balanece < amount {
		fmt.Printf("Tutar yetersiz Mevcut Bakiye : %v  , İstenilen Tutar : %v\n", w.Balanece, amount)
		return false, 0
	}
	w.Balanece -= amount
	return true, w.Balanece
}

func (w *Wallet) GetBalance() float64 {

	w.mtx.Lock()
	defer w.mtx.Unlock()
	return w.Balanece
}
