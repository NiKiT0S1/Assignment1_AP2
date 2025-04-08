package http

import (
	"github.com/gin-gonic/gin"
	"inventoryService/internal/domain"
)

func RegisterRoutes(r *gin.Engine, productUC domain.ProductUsecase) {
	h := &Handler{
		productUC: productUC,
		//categoryUC: categoryUC,
	}

	r.POST("/products", h.CreateProduct)
	r.GET("/products/:id", h.GetProduct)
	r.PATCH("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
	r.GET("/products", h.ListProducts)

	//r.POST("/categories", h.CreateCategory)
	//r.GET("/categories/:id", h.GetCategory)
	//r.PATCH("/categories/:id", h.UpdateCategory)
	//r.DELETE("/categories/:id", h.DeleteCategory)
	//r.GET("/categories", h.ListCategories)
}
