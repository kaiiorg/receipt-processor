package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kaiiorg/receipt-processor/internal/points_calculator"
)

type Api struct {
	calculator points_calculator.Calculator

	router *gin.Engine
}

func New() *Api {
	api := &Api{
		calculator: points_calculator.Calculator{},
		router:     gin.Default(),
	}

	api.router.POST("/receipts/process", api.processReceipt)
	api.router.GET("/receipts/:id/points")

	return api
}

func (api *Api) Run(port string) error {
	return api.router.Run(port)
}
