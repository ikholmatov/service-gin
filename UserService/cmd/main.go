package main

import (
	"github.com/venomuz/project4/UserService/config"
	pb "github.com/venomuz/project4/UserService/genproto"
	"github.com/venomuz/project4/UserService/pkg/db"
	"github.com/venomuz/project4/UserService/pkg/logger"
	"github.com/venomuz/project4/UserService/service"
	grpcClient "github.com/venomuz/project4/UserService/service/grpc_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	grocC, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("grpc connection to postservice error", logger.Error(err))
	}
	userService := service.NewUserService(connDB, log, grocC)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
