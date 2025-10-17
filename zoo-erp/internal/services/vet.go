package services

// VetClinic — проверка здоровья нового животного.
type VetClinic interface {
	CheckHealth(species string) bool
}

// SimpleVet — примитивная логика для ДЗ.
type SimpleVet struct{}

func NewSimpleVet() VetClinic { return &SimpleVet{} }

func (s *SimpleVet) CheckHealth(species string) bool {
	// Всё здорово, кроме искусственно помеченных "Sick..." видов.
	// Например, "SickRabbit" — будет отклонён.
	if len(species) >= 4 && species[:4] == "Sick" {
		return false
	}
	return true
}
