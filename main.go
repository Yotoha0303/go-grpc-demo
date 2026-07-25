package main

import (
	"log"

	"go-grpc-demo/config"
	"go-grpc-demo/pkg/db"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	if err := r.Run(":8083"); err != nil {
		log.Fatalf("http server: %v", err)
	}
}
