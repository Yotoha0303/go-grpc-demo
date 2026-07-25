package main

import (
	"fmt"
	"log"
	"net"

	pb "go-grpc-demo/api/proto/order"
	"go-grpc-demo/config"
	"go-grpc-demo/internal/grpc_server"
	"go-grpc-demo/internal/service"
	"go-grpc-demo/pkg/db"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("load env: %v", err)
	}

	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	gormDB, err := db.Init(&cfg.MySQL)
	if err != nil {
		log.Fatalf("init gorm: %v", err)
	}
	log.Printf("mysql connected: %s@%s:%d/%s",
		cfg.MySQL.User, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.Database)

	// reserved for dao / service wiring
	_ = gormDB

	orderSvc := service.NewOrderService(gormDB)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen gRPC:%v", err)
		}
		grpcServer := grpc.NewServer()

		pb.RegisterOrderServiceServer(grpcServer, grpc_server.NewOrderGRPCServer(orderSvc))

		fmt.Println("gRPC Server is running on :50051...")

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to server gRPC:%v", err)
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/api/v1/orders", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "HTTP order API"})
	})

	if err := r.Run(":8083"); err != nil {
		log.Fatalf("http server: %v", err)
	}
}
