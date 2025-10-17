package domain

import "errors"

var (
	ErrKindnessOutOfRange = errors.New("уровень доброты должен быть в [0..10]")
)

// Herbivore — травоядное: Animal + уровень доброты (0..10).
type Herbivore struct {
	Animal
	kindness int
}

func NewHerbivore(name string, number, food, kindness int) (*Herbivore, error) {
	if kindness < 0 || kindness > 10 {
		return nil, ErrKindnessOutOfRange
	}
	base, err := NewAnimal(name, number, food)
	if err != nil {
		return nil, err
	}
	return &Herbivore{Animal: *base, kindness: kindness}, nil
}

func (h Herbivore) Kindness() int { return h.kindness }

// Rabbit — конкретный вид (Herbivore).
type Rabbit struct{ Herbivore }

func NewRabbit(number, food, kindness int) (*Rabbit, error) {
	h, err := NewHerbivore("Rabbit", number, food, kindness)
	if err != nil {
		return nil, err
	}
	return &Rabbit{Herbivore: *h}, nil
}

// Monkey — упрощённо считаем травоядным для ТЗ.
type Monkey struct{ Herbivore }

func NewMonkey(number, food, kindness int) (*Monkey, error) {
	h, err := NewHerbivore("Monkey", number, food, kindness)
	if err != nil {
		return nil, err
	}
	return &Monkey{Herbivore: *h}, nil
}
