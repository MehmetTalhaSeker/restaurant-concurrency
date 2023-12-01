package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"sync"
	"time"
)

func orderWaiter(orders chan<- Order, wg *sync.WaitGroup, menu []Dish) {
	defer close(orders)
	defer wg.Done()

	dishNumbers := make([]int, 0, 1)

	for {
		var dishNumber int
		color.Green("\nSelect a dish (enter dish number or 'done' to finish): ")
		var input string
		fmt.Scan(&input)

		if input == "done" && len(dishNumbers) == 0 {
			color.Green("\nThe waiter thought you were mocking him and kicked you out of the restaurant.")
			os.Exit(0)
		} else if input == "done" {
			break
		}

		n, err := fmt.Sscanf(input, "%d", &dishNumber)
		if n != 1 || err != nil || dishNumber < 1 || dishNumber > len(menu) {
			color.Green("\n£-Invalid input. Please enter a valid dish number or 'done' to finish order.")
			continue
		}

		dishNumbers = append(dishNumbers, dishNumber)
	}

	color.Green("\norder waiter is going to the kitchen...")
	time.Sleep(time.Second * 4)

	for _, number := range dishNumbers {
		order := Order{Dish: menu[number-1]}
		orders <- order
	}
}

func serverWaiter(readyToServe chan Order, readyToPay chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(readyToServe)

	for {
		select {
		case order := <-readyToServe:
			color.Blue("\nServing waiter fetching the food...")

			time.Sleep(time.Second * 4)
			color.Blue("\nDish %s is ready! Bon appetit!", order.Dish.Name)

		case <-readyToPay:
			var input string
			color.Red("\n'payAndLeave' to leave the restaurant or try something else.")
			fmt.Scan(&input)

			if input == "payAndLeave" {
				os.Exit(0)
			} else {
				color.Red("\nso u choose leave the restaurant without paying")
				os.Exit(0)
			}
		}
	}
}
