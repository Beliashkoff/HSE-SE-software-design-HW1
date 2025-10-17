package repo

import (
	"testing"

	"zoo-erp/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestAnimalRepo_DuplicateNumber(t *testing.T) {
	r := NewAnimalMem()

	r1, _ := domain.NewRabbit(1, 3, 7)
	err := r.Add(r1)
	require.NoError(t, err)

	r2, _ := domain.NewRabbit(1, 4, 8) // тот же номер
	err = r.Add(r2)
	require.ErrorIs(t, err, ErrDuplicateNumber)

	all := r.All()
	require.Len(t, all, 1)
	require.Equal(t, 1, all[0].Number())
}

func TestAnimalRepo_AllReturnsCopy(t *testing.T) {
	r := NewAnimalMem()

	a1, _ := domain.NewRabbit(10, 2, 6)
	a2, _ := domain.NewRabbit(11, 3, 9)

	_ = r.Add(a1)
	_ = r.Add(a2)

	out := r.All()
	require.Len(t, out, 2)

	// Поменяем локально порядок в копии — репозиторий не должен измениться.
	out[0], out[1] = out[1], out[0]

	again := r.All()
	require.Equal(t, 10, again[0].Number())
	require.Equal(t, 11, again[1].Number())
}
