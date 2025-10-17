package repo

import (
	"sync"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/domain"
)

// animalRecord — узкий контракт для репозитория: нужно и есть, и быть инвентарным
type animalRecord interface {
	domain.Eater
	domain.InventoryItem
}

// AnimalRepo — интерфейс хранилища животных.
type AnimalRepo interface {
	Add(x animalRecord) error
	All() []animalRecord
	GetByNumber(n int) (animalRecord, bool)
}

type animalMem struct {
	mu   sync.RWMutex
	list []animalRecord // порядок добавления
	idx  map[int]int    // номер -> индекс в list
}

func NewAnimalMem() AnimalRepo {
	return &animalMem{
		idx: make(map[int]int),
	}
}

func (r *animalMem) Add(x animalRecord) error {
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

func (r *animalMem) All() []animalRecord {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]animalRecord, len(r.list))
	copy(out, r.list)
	return out
}

func (r *animalMem) GetByNumber(n int) (animalRecord, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	i, ok := r.idx[n]
	if !ok {
		return nil, false
	}
	return r.list[i], true
}
