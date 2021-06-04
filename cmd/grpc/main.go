package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pablocrivella/mancala/cmd/grpc/server"
	"github.com/pablocrivella/mancala/internal/engine"
	"github.com/pablocrivella/mancala/internal/postgres"
	pb "github.com/pablocrivella/mancala/pkg/proto/mancala/v1"
	"google.golang.org/grpc"
)

const exitFail = 1

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run() error {
	db, err := postgres.OpenDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	gameStore := postgres.NewGameStore(db)
	eng := engine.Service{
		GameStore: gameStore,
	}

	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGameServiceServer(grpcServer, server.NewGameService(eng))
	return grpcServer.Serve(lis)
}
