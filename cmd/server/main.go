package main

import (
	"log"

	"github.com/Wendiboy/users-service/internal/database"
	transportgrpc "github.com/Wendiboy/users-service/internal/transport/grpc"
	"github.com/Wendiboy/users-service/internal/user"
)

func main() {
	// database.InitDB()
	db, err := database.InitDB()

	if err != nil {
		log.Fatalf("Could not connect to DataBase: %v", err)
	}

	repo := user.NewRepository(db)
	svc := user.NewService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
