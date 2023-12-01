package main

import (
	"github.com/fatih/color"
	"sync"
	"time"
)

func kitchen(orders <-chan Order, readyToServe chan Order, readyToPay chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orders {
		color.Magenta("\n#-Preparing food:%s.", order.Dish.Name)
		time.Sleep(order.Dish.CookingTime)
		color.Magenta("\n#-Ding ding ding, %s is ready to serve.", order.Dish.Name)
		readyToServe <- order
	}
	readyToPay <- struct{}{}
}
