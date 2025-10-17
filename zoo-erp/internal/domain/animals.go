package domain

import (
	"errors"
	"strings"
)

type Animal struct {
	name            string
	inventoryNumber int
	dailyFoodKG     int
}

var (
	ErrEmptyName      = errors.New("имя не может быть пустым")
	ErrNegativeFood   = errors.New("Значение потребляемой еды не может быть отрицательным")
	ErrNegativeNumber = errors.New("Значение номера не может быть отрицательным")
)

func NewAnimal(name string, number, food int) (*Animal, error) {
	n := strings.TrimSpace(name)
	if food < 0 {
		return nil, ErrNegativeFood
	}
	if n == "" {
		return nil, ErrEmptyName
	}
	if number < 0 {
		return nil, ErrNegativeNumber
	}
	return &Animal{
		name:            n,
		inventoryNumber: number,
		dailyFoodKG:     food,
	}, nil
}

func (a Animal) Food() int {
	return a.dailyFoodKG
}

func (a Animal) Name() string {
	return a.name
}

func (a Animal) Number() int {
	return a.inventoryNumber
}
