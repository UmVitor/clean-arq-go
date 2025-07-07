package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) GetAllOrders() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	orderOutputDTO := []OrderOutputDTO{}
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		orderOutputDTO = append(orderOutputDTO, dto)
	}

	return orderOutputDTO, nil
}
