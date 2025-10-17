package repo

import (
	"sync"

	"zoo-erp/internal/domain"
)

// ThingRepo — хранилище инвентарных вещей (не животных).
type ThingRepo interface {
	Add(x domain.InventoryItem) error
	All() []domain.InventoryItem
}

type thingMem struct {
	mu   sync.RWMutex
	list []domain.InventoryItem
	idx  map[int]int
}

func NewThingMem() ThingRepo {
	return &thingMem{
		idx: make(map[int]int),
	}
}

func (r *thingMem) Add(x domain.InventoryItem) error {
	n := x.Number()

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.idx[n]; exists {
		return ErrDuplicateNumber
	}

	r.list = append(r.list, x)
	r.idx[n] = len(r.list) - 1
	return nil
}

func (r *thingMem) All() []domain.InventoryItem {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]domain.InventoryItem, len(r.list))
	copy(out, r.list)
	return out
}
