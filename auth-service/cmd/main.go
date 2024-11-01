package main

import (
	"fmt"

	"github.com/lunarKettle/task-management-platform/auth-service/internal/infrastructure"
	"github.com/lunarKettle/task-management-platform/auth-service/internal/transport"
	"github.com/lunarKettle/task-management-platform/auth-service/internal/usecases"

	"github.com/joho/godotenv"

	"log"
	"os"
)

const (
	ServerAddressEnv    = "SERVER_ADDRESS"
	ConnectionStringEnv = "CONNECTION_STRING"
	SecretKeyEnv        = "SECRET_KEY"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	database := infrastructure.NewDatabase()
	err := database.OpenConnetion(os.Getenv(ConnectionStringEnv))
	if err != nil {
		return err
	}
	repo := infrastructure.NewUserRepository(database)
	jwtGenerator := infrastructure.NewJWTGenerator(os.Getenv(SecretKeyEnv))
	authUseCases := usecases.NewAuthUseCases(repo, jwtGenerator)
	server := transport.NewGRPCServer(authUseCases)
	err = server.Start(os.Getenv(ServerAddressEnv))
	if err != nil {
		return err
	}
	return nil
}
