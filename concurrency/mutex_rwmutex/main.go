package main

import (
	"awesomeProject/concurrency/mutex_rwmutex/domain"
	"awesomeProject/concurrency/mutex_rwmutex/worker"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Structs

func main() {
	var wg sync.WaitGroup
	orderChan := make(chan domain.Order)
	//rwmtx := sync.RWMutex{}
	//mtx := sync.Mutex{}
	prices := map[string]float64{"BTC": 85, "ETH": 50, "DOGE": 80}
	market := domain.NewMarket(prices)
	wallet := domain.NewWallet(1200)

	go func() {
		for {

			market.UpdatePrice("BTC", float64(rand.IntN(90)+40))
			time.Sleep(130 * time.Millisecond)
			fmt.Println(market.GetPrice("BTC"))
		}
	}()

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go worker.Run(i, orderChan, wallet, market, &wg)
	}

	for i := 0; i < 10; i++ {
		orderChan <- domain.Order{
			Symbol: "BTC",
			Amount: rand.Float64() * 2,
			Type:   domain.Buy,
		}
	}

	close(orderChan)
	wg.Wait()

	fmt.Printf("\n--- SISTEM KAPANDI ---\nKalan Bakiye: %.2f USD\n", wallet.GetBalance())
}
