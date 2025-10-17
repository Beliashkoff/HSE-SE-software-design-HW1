package app

import (
	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/services"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/repo"

	"go.uber.org/dig"
)

// Build — composition root: регистрируем все зависимости.
func Build() *dig.Container {
	c := dig.New()

	// Репозитории (in-memory)
	_ = c.Provide(repo.NewAnimalMem)
	_ = c.Provide(repo.NewThingMem)

	// Сервисы и зависимости
	_ = c.Provide(services.NewSimpleVet)
	_ = c.Provide(services.NewSeq)
	_ = c.Provide(services.NewZoo)

	return c
}
