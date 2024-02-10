package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) FindAllOrder(ctx context.Context, in *pb.Empty) (*pb.FindAllOrdersResponse, error) {
	dto := usecase.ListOrderInputDTO{}
	output, err := s.ListOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	var list []*pb.Order
	for _, data := range output.Orders {
		list = append(list, &pb.Order{
			Id:         data.ID,
			Price:      float32(data.Price),
			Tax:        float32(data.Tax),
			FinalPrice: float32(data.FinalPrice),
		})
	}
	return &pb.FindAllOrdersResponse{
		Orders: list,
	}, nil
}
