package memory_repository

import (
	"os"
	"sync"

	"github.com/kaiiorg/receipt-processor/internal/models"
)

type MemoryRepository struct {
	store sync.Map
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (mr *MemoryRepository) SaveReceipt(id string, receipt *models.Receipt) error {
	mr.store.Store(id, receipt)
	return nil
}

func (mr *MemoryRepository) LoadReceipt(id string) (*models.Receipt, error) {
	receipt, ok := mr.store.Load(id)
	if !ok {
		return nil, os.ErrNotExist
	}

	return receipt.(*models.Receipt), nil
}
