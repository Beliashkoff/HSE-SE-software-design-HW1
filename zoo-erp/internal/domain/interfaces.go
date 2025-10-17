package domain

// Eater - килограммы в сутки
type Eater interface {
	Food() int
}

// InventoryItem - инвентарный номер по счёту
type InventoryItem interface {
	Number() int
	Name() string
}
