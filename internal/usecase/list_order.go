package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrderInputDTO struct {
}
type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderOutputDTO struct {
	Orders []Order
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute(input ListOrderInputDTO) (ListOrderOutputDTO, error) {
	var ordersList []Order

	orders, err := l.OrderRepository.FindAll()

	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	for _, order := range orders {
		ordersList = append(ordersList, Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	var dto ListOrderOutputDTO = ListOrderOutputDTO{
		Orders: ordersList,
	}

	l.OrderListed.SetPayload(dto)
	l.EventDispatcher.Dispatch(l.OrderListed)

	return dto, nil
}
