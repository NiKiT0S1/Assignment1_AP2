// package cmd
package main

import (
	"inventoryService/internal/delivery/http"
	"inventoryService/internal/repository"
	"inventoryService/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect DB:", err)
	}

	productRepo := repository.NewProductRepo(db)
	//categoryRepo := repository.NewCategoryRepo(db)

	productUC := usecase.NewProductUsecase(productRepo)
	//categoryUC := usecase.NewCategoryUsecase(categoryRepo)

	router := gin.Default()
	//http.RegisterRoutes(router, productUC, categoryUC)
	http.RegisterRoutes(router, productUC)

	router.Run(":8081")
}
