package main

import (
	"mushrooms-api/controller"
	"mushrooms-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductUsecase := usecase.NewProductUsecase()

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")

}
