package repository

import "github.com/kaiiorg/receipt-processor/internal/models"

type Repository interface {
	SaveReceipt(id string, receipt *models.Receipt) error
	LoadReceipt(id string) (*models.Receipt, error)
}
