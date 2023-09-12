package main

import "github.com/mario-aj/go-intesivo/internal/entity"

func main() {
	order, err := entity.NewOrder("123", 10, 2)

	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}	
}
