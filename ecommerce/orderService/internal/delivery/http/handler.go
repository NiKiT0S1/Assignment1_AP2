package http

import (
	"net/http"
	"orderService/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc domain.OrderUsecase
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *Handler) GetOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.UpdateStatus(id, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) ListOrdersByUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	orders, err := h.uc.ListByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
