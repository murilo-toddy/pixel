package grpc

import (
	"context"

	"github.com/murilo-toddy/pixel/app/grpc/pb"
	"github.com/murilo-toddy/pixel/app/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func (s *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := s.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (s *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	key, err := s.PixUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:        key.ID,
		Kind:      key.Kind,
		Key:       key.Key,
		CreatedAt: key.Account.CreatedAt.String(),
		Account: &pb.Account{
			AccountId:     key.AccountID,
			AccountNumber: key.Account.Number,
			BankId:        key.Account.BankID,
			BankName:      key.Account.Bank.Name,
			OwnerName:     key.Account.OwnerName,
			CreatedAt:     key.Account.CreatedAt.String(),
		},
	}, nil
}

func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}
