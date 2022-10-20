package diccionario_test

import (
	dic "diccionario/abb"
	"diccionario/errores"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAbb_Guardar(t *testing.T) {
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
	abb.Guardar(3, 2)
	require.EqualValues(t, 2, abb.Obtener(3))
	require.True(t, abb.Pertenece(1))
	require.PanicsWithValue(t, errores.ErrorNoEncontrado{}.Error(), func() { abb.Obtener(5) })
	require.PanicsWithValue(t, errores.ErrorNoEncontrado{}.Error(), func() { abb.Borrar(5) })
	abb.Guardar(4, 2)
	abb.Guardar(5, 2)
	require.EqualValues(t, 2, abb.Borrar(3))
}

func TestAthus(t *testing.T) {
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
	abb.Guardar(2, 2)
	abb.Guardar(4, 4)
	abb.Borrar(3)
}
