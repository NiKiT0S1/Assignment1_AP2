// package cmd
package main

import (
	"log"
	"orderService/internal/delivery/http"
	"orderService/internal/repository"
	"orderService/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln("DB connection failed:", err)
	}

	orderRepo := repository.NewOrderRepo(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)

	r := gin.Default()
	http.RegisterRoutes(r, orderUC)

	r.Run(":8082")
}
