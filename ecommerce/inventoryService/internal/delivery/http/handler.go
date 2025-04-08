package http

import (
	"inventoryService/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	productUC domain.ProductUsecase
	//categoryUC domain.CategoryUsecase
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.productUC.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *Handler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.productUC.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.ID = id
	if err := h.productUC.Update(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.productUC.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) ListProducts(c *gin.Context) {
	products, err := h.productUC.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
