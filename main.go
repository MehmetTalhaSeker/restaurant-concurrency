package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	menu := []Dish{
		{Name: "Burger", CookingTime: 6 * time.Second},
		{Name: "Salad", CookingTime: 1 * time.Second},
		{Name: "Spaghetti", CookingTime: 3 * time.Second},
		{Name: "Pizza", CookingTime: 10 * time.Second},
	}

	fmt.Println("Menu:")
	for i, dish := range menu {
		fmt.Printf("%d. %s (%s)\n", i+1, dish.Name, dish.CookingTime)
	}

	var wg sync.WaitGroup

	orders := make(chan Order)
	readyToServe := make(chan Order)
	readyToPay := make(chan struct{})

	wg.Add(1)
	go kitchen(orders, readyToServe, readyToPay, &wg)

	wg.Add(1)
	go orderWaiter(orders, &wg, menu)

	wg.Add(1)
	go serverWaiter(readyToServe, readyToPay, &wg)

	wg.Wait()
}
