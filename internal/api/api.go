package api

import (
	"github.com/kaiiorg/receipt-processor/internal/points_calculator"
	"github.com/kaiiorg/receipt-processor/internal/repository"

	"github.com/gin-gonic/gin"
)

type Api struct {
	calculator points_calculator.Calculator
	repo       repository.Repository

	router *gin.Engine
}

func New(repo repository.Repository) *Api {
	api := &Api{
		calculator: points_calculator.Calculator{},
		repo:       repo,
		router:     gin.Default(),
	}

	api.router.POST("/receipts/process", api.processReceipt)
	api.router.GET("/receipts/:id/points", api.points)

	return api
}

func (api *Api) Run(port string) error {
	return api.router.Run(port)
}
