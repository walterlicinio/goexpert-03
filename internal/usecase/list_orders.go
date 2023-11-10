package usecase

import (
	"github.com/walterlicinio/goexpert-03/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]entity.Order, error) {
	orders, err := l.OrderRepository.ListAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
