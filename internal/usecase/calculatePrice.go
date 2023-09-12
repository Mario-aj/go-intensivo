package usecase

import "github.com/Mario-aj/go-intensivo/internal/entity"

type OrderInput struct {
	ID    string
	price float64
	tax   float64
}

type OrderOutput struct {
	ID         string
	price      float64
	tax        float64
	finalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrice)  Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.price, input.tax)

	if err != nil {
		return  nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
			return nil, err
	}

	return &OrderOutput{
		ID: order.ID,
		tax: order.Tax,
		price: order.Price,
		finalPrice: order.FinalPrice,
	}, nil
}
