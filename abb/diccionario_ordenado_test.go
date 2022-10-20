package diccionario_test

import (
	dic "diccionario/abb"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAbb_Guardar(t *testing.T) {
	require.True(t, true)
	abb := dic.CrearABB[int, int](func(x int, y int) int {
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	})
	abb.Guardar(1, 1)
	abb.Guardar(3, 3)
	abb.Borrar(3)
}
