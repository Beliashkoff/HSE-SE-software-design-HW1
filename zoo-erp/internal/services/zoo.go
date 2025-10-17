package services

import (
	"zoo-erp/internal/domain"
	"zoo-erp/internal/repo"
)

// ZooService — сценарии зоопарка.
type ZooService interface {
	AcceptAnimal(x domain.InventoryItem) error
	AddThing(x domain.InventoryItem) error
	TotalFoodKG() int
	ContactZoo() []domain.InventoryItem
	Inventory() []domain.InventoryItem
}

type zoo struct {
	animals repo.AnimalRepo
	things  repo.ThingRepo
	vet     VetClinic
}

func NewZoo(animals repo.AnimalRepo, things repo.ThingRepo, vet VetClinic) ZooService {
	return &zoo{animals: animals, things: things, vet: vet}
}

func (z *zoo) AcceptAnimal(x domain.InventoryItem) error {
	// Проверка здоровья по виду/наименованию.
	if !z.vet.CheckHealth(x.Name()) {
		return ErrRejectedByVet
	}
	// Репозиторий гарантирует, что x действительно "животное" (Eater), и соблюдает уникальность номера.
	if err := z.animals.Add(x); err != nil {
		return err
	}
	return nil
}

func (z *zoo) AddThing(x domain.InventoryItem) error {
	return z.things.Add(x)
}

func (z *zoo) TotalFoodKG() int {
	sum := 0
	for _, it := range z.animals.All() {
		if e, ok := it.(domain.Eater); ok {
			sum += e.Food()
		}
	}
	return sum
}

func (z *zoo) ContactZoo() []domain.InventoryItem {
	type kinder interface{ Kindness() int }
	var res []domain.InventoryItem
	for _, it := range z.animals.All() {
		if h, ok := it.(kinder); ok && h.Kindness() > 5 {
			res = append(res, it)
		}
	}
	return res
}

func (z *zoo) Inventory() []domain.InventoryItem {
	// Животные + вещи
	animals := z.animals.All()
	things := z.things.All()
	res := make([]domain.InventoryItem, 0, len(animals)+len(things))
	res = append(res, animals...)
	res = append(res, things...)
	return res
}
