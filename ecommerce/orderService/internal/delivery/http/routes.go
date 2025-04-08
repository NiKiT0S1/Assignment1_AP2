package http

import (
	"orderService/internal/domain"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, uc domain.OrderUsecase) {
	h := &Handler{uc}

	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("/orders/:id", h.UpdateOrderStatus)
	r.GET("/orders", h.ListOrdersByUser)
}
