package main

import "time"

type Dish struct {
	Name        string
	CookingTime time.Duration
}

type Order struct {
	Dish Dish
}
