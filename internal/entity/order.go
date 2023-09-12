package entity

import "errors"

type Order struct {
	ID         string
	price      float64
	tax        float64
	finalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		price: price,
		tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) Validate() error {
	if order.ID == "" {
		return errors.New("id is required")
	}

	if order.price <= 0 {
		return errors.New("price must be grater than zero")
	}

	if order.tax < 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	err := order.Validate()

	if err != nil {
		return err
	}

	order.finalPrice = order.price + order.tax

	return nil
}
