package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
	"zoo-erp/internal/domain"
)

func TestThingRepo_DuplicateNumber(t *testing.T) {
	r := NewThingMem()

	t1, _ := domain.NewTable(5)
	err := r.Add(t1)
	require.NoError(t, err)

	t2, _ := domain.NewComputer(5) // тот же номер
	err = r.Add(t2)
	require.ErrorIs(t, err, ErrDuplicateNumber)

	all := r.All()
	require.Len(t, all, 1)
	require.Equal(t, 5, all[0].Number())
}
