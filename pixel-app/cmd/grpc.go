package cmd

import (
	"os"

	"github.com/murilo-toddy/pixel/app/grpc"
	"github.com/murilo-toddy/pixel/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
        database := db.Connect(os.Getenv("env"))
        grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
    rootCmd.PersistentFlags().IntVarP(&portNumber, "port", "p", 50051, "gRPC server port")
    rootCmd.AddCommand(grpcCmd)
}

