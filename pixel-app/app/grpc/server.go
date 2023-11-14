package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/murilo-toddy/pixel/app/grpc/pb"
	"github.com/murilo-toddy/pixel/app/usecase"
	"github.com/murilo-toddy/pixel/domain/infrastructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
    grpcServer := grpc.NewServer()
    reflection.Register(grpcServer)

    pixRepository := repository.PixKeyRepositoryDB{DB: database}
    pixUseCase := usecase.PixUseCase{PixKeyRepository: &pixRepository}
    pixGrpcService := NewPixGrpcService(pixUseCase)
    pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

    address := fmt.Sprintf("0.0.0.0:%d", port)
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatal("Could not start grpc server: ", err)
    }

    log.Printf("gRPC server has started on port %d", port)
    err = grpcServer.Serve(listener)
    if err != nil {
        log.Fatal("gRPC server stopped: ", err)
    }
}
