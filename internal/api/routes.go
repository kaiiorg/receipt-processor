package api

import (
	"fmt"
	"net/http"

	"github.com/kaiiorg/receipt-processor/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (api *Api) processReceipt(c *gin.Context) {
	// Pull the receipt from the body
	r := &models.Receipt{}
	err := c.BindJSON(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate
	err = r.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Persist
	id := uuid.NewString()

	// Respond
	c.JSON(200, gin.H{"id": id})
}

func (api *Api) points(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	// Pull persisted receipt
	fmt.Sprintf("id: %s", id) // NOOP for now

	// Calculate the points

	// Return the result
	c.JSON(200, gin.H{"points": 0})
}
