package repo

import "errors"

var (
	ErrDuplicateNumber = errors.New("duplicate inventory number")
	ErrNotAnimal       = errors.New("item is not an animal (no Eater role)")
)
