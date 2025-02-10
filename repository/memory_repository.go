package repository

import (
	"fetch-assessment/rules"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// an interface for an in-memory repository
type MemoryRepository struct {
	mutex    sync.RWMutex                         // for thread safety
	receipts map[string] /* uuid */ rules.Receipt // just use the business object since there isn't a clear need for a data entity and we've already showcased the usage/strength of using mappers to convert between types
}

func NewMemoryRepository() *MemoryRepository {
	// initialize and return an in-memory repository
	return &MemoryRepository{
		receipts: make(map[string]rules.Receipt),
	}
}

// persist the given receipt to the in-memory repository
func (r *MemoryRepository) SaveReceipt(receipt rules.Receipt) (string, error) {
	r.mutex.Lock()         // obtain a WRITE lock to ensure thread safety
	defer r.mutex.Unlock() // ensure we release our WRITE lock upon the close of this scope

	id := uuid.New().String() // grab a new UUID
	r.receipts[id] = receipt  // 'save' the receipt to the repository
	return id, nil
}

// load the receipt specified by the given uuid (string) from the in-memory repository
func (r *MemoryRepository) LoadReceipt(id string) (rules.Receipt, error) {
	r.mutex.RLock()         // obtain a non-exclusive READ lock
	defer r.mutex.RUnlock() // ensure we release our READ lock so as not to block any writes

	// grab the receipt from the hashtable if it exists
	receipt, exists := r.receipts[id]
	if !exists {
		// return an empty object... (change method signature to *rules.Receipt and return nil instead, possibly??)
		return rules.Receipt{}, fmt.Errorf("receipt with ID %s was not found in the repository", id)
	}
	return receipt, nil
}
