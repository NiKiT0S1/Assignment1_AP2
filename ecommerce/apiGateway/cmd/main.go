// package cmd
package main

import (
	"apiGateway/internal/middleware"
	_ "net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.BasicAuth())

	// Reverse proxies
	inventoryURL, _ := url.Parse("http://localhost:8081")
	orderURL, _ := url.Parse("http://localhost:8082")

	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)

	// Route delegation
	r.Any("/products", gin.WrapH(inventoryProxy))
	r.Any("/products/*proxyPath", gin.WrapH(inventoryProxy))

	r.Any("/orders", gin.WrapH(orderProxy))
	r.Any("/orders/*proxyPath", gin.WrapH(orderProxy))

	r.Run(":8080")
}
