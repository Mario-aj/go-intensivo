package entity

type Order struct {
	ID         string
	price      float64
	tax        float64
	finalPrice float64
}

func NewOrder(id string, price float64, tax float64) *Order {
	return &Order{
		ID:    id,
		price: price,
		tax:   tax,
	}
}

func (order *Order) CalculateFinalPrice() {
	order.finalPrice = order.price + order.tax
}
