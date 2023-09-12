package main

import "fmt"

type Car struct {
	model string
	color string
}

func (c *Car) setColor(color string) { // point to the same memory spot
	c.color = color
}

func (c Car) setModel(model string) { // Not point to the same momery spot.
	println("The model that you want to set is", model)
}

func sum(x, y int) int {
	return x + y
}

func main() {
	println("hello wolrd!")

	car := Car{
		model: "Ferrari",
		color: "White",
	}

	println(fmt.Sprintf("I have a nice car of model %s and the color %s", car.model, car.color))
	println("the sum of 6 with 5 is:", sum(6, 5))

	car.setModel("Hyunday")
	car.setColor("Red")

	println(fmt.Sprintf("I have a nice car of model %s and the color %s", car.model, car.color))

	a := 10
	b := &a

	*b = 6

	println(a)
	println(b)
}
