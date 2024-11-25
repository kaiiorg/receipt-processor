package api

import (
	"errors"
	"net/http"
	"os"

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
	err = api.repo.SaveReceipt(id, r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	r, err := api.repo.LoadReceipt(id)
	if errors.Is(err, os.ErrNotExist) {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if r == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "saved receipt was nil!"})
		return
	}

	// Calculate the points
	points := api.calculator.Calculate(*r)

	// Return the result
	c.JSON(200, gin.H{"points": points})
}
