package domain

// Predator — хищник: Animal без доброты.
type Predator struct {
	Animal
}

func NewPredator(name string, number, food int) (*Predator, error) {
	base, err := NewAnimal(name, number, food)
	if err != nil {
		return nil, err
	}
	return &Predator{Animal: *base}, nil
}

type Tiger struct{ Predator }

func NewTiger(number, food int) (*Tiger, error) {
	p, err := NewPredator("Tiger", number, food)
	if err != nil {
		return nil, err
	}
	return &Tiger{Predator: *p}, nil
}

type Wolf struct{ Predator }

func NewWolf(number, food int) (*Wolf, error) {
	p, err := NewPredator("Wolf", number, food)
	if err != nil {
		return nil, err
	}
	return &Wolf{Predator: *p}, nil
}
