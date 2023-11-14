package main

import (
	"os"

	"github.com/murilo-toddy/pixel/app/grpc"
	"github.com/murilo-toddy/pixel/infrastructure/db"
)

func main() {
    database := db.Connect(os.Getenv("env"))
    grpc.StartGrpcServer(database, 50051)
}
