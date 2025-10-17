package domain

import "fmt"

type Animal struct {
	type_of_animal   string
	inventory_number int
	daily_ration     int
}

func NewAnimal(name string, number, food int) (*Animal, error) {
	if food < 0 {
		return nil, fmt.Errorf("Потребление меньше 0 не может быть")
	}
	if name == "" {
		return nil, fmt.Errorf("Имя не может быть пустым")
	}
	return &Animal{
		type_of_animal:   name,
		inventory_number: number,
		daily_ration:     food,
	}, nil
}

func (a Animal) Food() int {
	return a.daily_ration
}

func (a Animal) Name() string {
	return a.type_of_animal
}

func (a Animal) Number() int {
	return a.inventory_number
}
