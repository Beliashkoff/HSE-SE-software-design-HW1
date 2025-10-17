package repo

import (
	"sync"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/domain"
)

// AnimalRepo — хранилище животных (только сущности, у которых есть Food()).
type AnimalRepo interface {
	Add(x domain.InventoryItem) error
	All() []domain.InventoryItem
	GetByNumber(n int) (domain.InventoryItem, bool)
}

type animalMem struct {
	mu   sync.RWMutex
	list []domain.InventoryItem // порядок добавления
	idx  map[int]int            // номер -> индекс в list
}

func NewAnimalMem() AnimalRepo {
	return &animalMem{
		idx: make(map[int]int),
	}
}

func (r *animalMem) Add(x domain.InventoryItem) error {
	// Защитимся от «вещей»: животное обязано быть Eater.
	if _, ok := x.(domain.Eater); !ok {
		return ErrNotAnimal
	}

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

func (r *animalMem) All() []domain.InventoryItem {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]domain.InventoryItem, len(r.list))
	copy(out, r.list)
	return out
}

func (r *animalMem) GetByNumber(n int) (domain.InventoryItem, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	i, ok := r.idx[n]
	if !ok {
		return nil, false
	}
	return r.list[i], true
}
