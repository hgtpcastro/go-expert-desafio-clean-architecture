package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type GetOrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (g *GetOrdersUseCase) Execute(page, limit int, sort string) (GetOrdersOutputDTO, error) {

	orders, err := g.OrderRepository.FindAll(page, limit, sort)
	if err != nil {
		return GetOrdersOutputDTO{}, err
	}

	var ordersOutputDTO []OrderOutputDTO

	for _, order := range orders {
		orderOutputDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersOutputDTO = append(ordersOutputDTO, orderOutputDTO)
	}

	return GetOrdersOutputDTO{Orders: ordersOutputDTO}, nil
}
