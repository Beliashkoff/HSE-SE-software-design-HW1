package services

import (
	"testing"

	"github.com/stretchr/testify/require"
	"zoo-erp/internal/domain"
	"zoo-erp/internal/repo"
)

type vetMock struct{ ok bool }

func (v vetMock) CheckHealth(string) bool { return v.ok }

func TestAcceptAnimal_RejectedByVet(t *testing.T) {
	z := NewZoo(repo.NewAnimalMem(), repo.NewThingMem(), vetMock{ok: false})

	r, _ := domain.NewRabbit(1, 2, 7)
	err := z.AcceptAnimal(r)
	require.ErrorIs(t, err, ErrRejectedByVet)
}

func TestTotalFood_Sum(t *testing.T) {
	z := NewZoo(repo.NewAnimalMem(), repo.NewThingMem(), vetMock{ok: true})

	r1, _ := domain.NewRabbit(1, 2, 7)
	r2, _ := domain.NewRabbit(2, 3, 4)

	require.NoError(t, z.AcceptAnimal(r1))
	require.NoError(t, z.AcceptAnimal(r2))

	require.Equal(t, 5, z.TotalFoodKG())
}

func TestContactZoo_FilterByKindness(t *testing.T) {
	z := NewZoo(repo.NewAnimalMem(), repo.NewThingMem(), vetMock{ok: true})

	r1, _ := domain.NewRabbit(1, 2, 6) // >5
	r2, _ := domain.NewMonkey(2, 2, 10)
	tg, _ := domain.NewTiger(3, 7)     // не травоядное

	_ = z.AcceptAnimal(r1)
	_ = z.AcceptAnimal(r2)
	_ = z.AcceptAnimal(tg)

	list := z.ContactZoo()
	require.Len(t, list, 2)
	names := []string{list[0].Name(), list[1].Name()}
	require.ElementsMatch(t, []string{"Rabbit", "Monkey"}, names)
}
