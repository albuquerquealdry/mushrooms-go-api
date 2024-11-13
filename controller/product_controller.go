package controller

import (
	"mushrooms-api/model"
	"mushrooms-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "teste",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
