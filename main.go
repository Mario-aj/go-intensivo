package main

import "github.com/Mario-aj/go-intensivo/internal/entity"

func main() {
	order, err := entity.NewOrder("123", 10, 2)

	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}
}
