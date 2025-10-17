package domain

import (
	"errors"
	"strings"
)

var (
	ErrThingEmptyName         = errors.New("вещь не может быть пуста")
	ErrThingNegativeInvNumber = errors.New("номер вещи не может быть отрицательным")
)

// Thing — базовый инвентарный предмет.
type Thing struct {
	name            string
	inventoryNumber int
}

func NewThing(name string, number int) (*Thing, error) {
	n := strings.TrimSpace(name)
	if n == "" {
		return nil, ErrThingEmptyName
	}
	if number < 0 {
		return nil, ErrThingNegativeInvNumber
	}
	return &Thing{name: n, inventoryNumber: number}, nil
}

func (t Thing) Name() string { return t.name }
func (t Thing) Number() int  { return t.inventoryNumber }

// Table — стол.
type Table struct{ Thing }

func NewTable(number int) (*Table, error) {
	th, err := NewThing("Table", number)
	if err != nil {
		return nil, err
	}
	return &Table{Thing: *th}, nil
}

// Computer — компьютер.
type Computer struct{ Thing }

func NewComputer(number int) (*Computer, error) {
	th, err := NewThing("Computer", number)
	if err != nil {
		return nil, err
	}
	return &Computer{Thing: *th}, nil
}
