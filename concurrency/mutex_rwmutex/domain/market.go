package domain

import "sync"

type Market struct {
	Prices map[string]float64
	mtx    sync.RWMutex
}

func NewMarket(priceMap map[string]float64) *Market {
	return &Market{
		Prices: priceMap,
	}
}

func (m *Market) UpdatePrice(symbol string, newPrice float64) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.Prices[symbol] = newPrice
}

func (m *Market) GetPrice(symbol string) float64 {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.Prices[symbol]
}
