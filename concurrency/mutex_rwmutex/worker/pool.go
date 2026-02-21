package worker

import (
	"awesomeProject/concurrency/mutex_rwmutex/domain"
	"fmt"
	"sync"
	"time"
)

func Run(id int, orders <-chan domain.Order, w *domain.Wallet, m *domain.Market, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d. Worker \n", id)
	for p := range orders {
		price := m.GetPrice(p.Symbol)
		if p.Type == domain.OrderType(domain.Buy) {

			ok, fl := w.ProcessBuy(price * p.Amount)
			if ok {
				fmt.Printf("Worker %d: ALIM BASARILI | Tutar: %.2f   || Bakiye : %f\n", id, price*p.Amount, fl)

			}
			time.Sleep(150 * time.Millisecond)
		}
	}
}
