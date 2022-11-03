package cola_prioridad_test

import (
	TDAheap "cola_prioridad"
	"cola_prioridad/errores"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHeapVacio(t *testing.T) {
	t.Log("probamos crear un heap vacio y ver si los errores funcionan adecuadamente")
	h := TDAheap.CrearHeap(strings.Compare)
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
	require.PanicsWithValue(t, errores.ErrorColaVacia{}.Error(), func() {
		h.Desencolar()
	})
	require.PanicsWithValue(t, errores.ErrorColaVacia{}.Error(), func() {
		h.VerMax()
	})
}

func TestHeapEncolarYDesencolarUnDato(t *testing.T) {
	t.Log("Encolamos y luego desencolamos un string")
	h := TDAheap.CrearHeap[string](strings.Compare)
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
	h.Encolar("hola")
	require.EqualValues(t, "hola", h.VerMax())
	require.EqualValues(t, 1, h.Cantidad())
	require.False(t, h.EstaVacia())
	require.EqualValues(t, "hola", h.Desencolar())
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapEncolarYDesencolarAlgunosDatos(t *testing.T) {
	t.Log("Encolamos y luego desencolamos algunos strings")
	h := TDAheap.CrearHeap[string](strings.Compare)
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
	letras := []string{"a", "b", "c", "d"}
	for cant, letra := range letras {
		h.Encolar(letra)
		require.EqualValues(t, letra, h.VerMax())
		require.EqualValues(t, cant+1, h.Cantidad())
	}

	for i := 1; i < len(letras)+1; i++ {
		require.EqualValues(t, letras[len(letras)-i], h.Desencolar())
	}
	require.True(t, h.EstaVacia())
}

func TestCrearHeapDesdeArray(t *testing.T) {
	t.Log("Creamos un Heap con datos insertados predeterminadamente")
	letras := []string{"a", "b", "c", "d"}
	h := TDAheap.CrearHeapArr[string](letras, strings.Compare)
	for i := 1; i < len(letras)+1; i++ {
		require.EqualValues(t, letras[len(letras)-i], h.Desencolar())
	}
	h.Encolar("e")
	h.Encolar("f")
	require.EqualValues(t, "f", h.Desencolar())
	require.EqualValues(t, "e", h.Desencolar())
	require.True(t, h.EstaVacia())
}

func TestCrearHeapDesdeArrayVacio(t *testing.T) {
	t.Log("Creamos un Heap con un array vacio")
	letras := []string{}
	h := TDAheap.CrearHeapArr[string](letras, strings.Compare)
	h.Encolar("e")
	h.Encolar("f")
	require.EqualValues(t, "f", h.Desencolar())
	require.EqualValues(t, "e", h.Desencolar())
	require.True(t, h.EstaVacia())
}

func TestVolumen(t *testing.T) {
	t.Log("Prueba con un gran volumen de datos")
	const _VOLUMEN = 10000
	h := TDAheap.CrearHeap[int](func(x, y int) int {
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	})
	for i := 0; i <= _VOLUMEN; i++ {
		h.Encolar(i)
		require.EqualValues(t, i+1, h.Cantidad())
	}
	require.EqualValues(t, _VOLUMEN, h.VerMax())
	for i := _VOLUMEN; i >= 0; i-- {
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestVolumenDesdeArray(t *testing.T) {
	t.Log("Prueba con un gran volumen de datos desde un arreglo muy grande")
	const _VOLUMEN = 1000
	arrayGrande := make([]int, _VOLUMEN)
	for i := 0; i < _VOLUMEN; i++ {
		arrayGrande[i] = i
	}
	h := TDAheap.CrearHeapArr[int](arrayGrande, func(x, y int) int {
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	})

	require.EqualValues(t, _VOLUMEN-1, h.VerMax())
	for i := _VOLUMEN - 1; i >= 0; i-- {
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestBorde(t *testing.T) {
	t.Log("Pruebas de caso borde")
	const _MUCHOS = 100
	h := TDAheap.CrearHeap[int](func(x, y int) int {
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	})
	for i := 0; i <= _MUCHOS; i++ {
		h.Encolar(i)
		require.EqualValues(t, i+1, h.Cantidad())
	}

	for i := 0; i < 100; i++ {
		h.Encolar(101)
		require.EqualValues(t, 101, h.Desencolar())
	}

	for i := _MUCHOS; i >= 0; i-- {
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i, h.Cantidad())
	}

	for i := 1; i <= 5; i++ {
		h.Encolar(i)
	}
	require.EqualValues(t, 5, h.Cantidad())
	require.EqualValues(t, 5, h.Desencolar())
	require.EqualValues(t, 4, h.Cantidad())
	require.False(t, h.EstaVacia())

}

func TestStuctHeap(t *testing.T) {
	t.Log("Creamos un heap de struct con su propio metodo de comparacion")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}
	f := func(a1, a2 avanzado) int {
		return strings.Compare(a1.x.a, a2.x.a)
	}

	a1 := avanzado{w: 10, z: "que", x: basico{a: "onda", b: 8}, y: basico{a: "gente", b: 10}}
	a2 := avanzado{w: 10, z: "piola", x: basico{a: "ahre", b: 14}, y: basico{a: "epico", b: 5}}
	a3 := avanzado{w: 10, z: "sant1", x: basico{a: "max0", b: 8}, y: basico{a: "cracks", b: 4}}

	h := TDAheap.CrearHeap[avanzado](f)
	h.Encolar(a1)
	h.Encolar(a2)
	h.Encolar(a3)
	require.EqualValues(t, 3, h.Cantidad())

	require.EqualValues(t, a1, h.VerMax())
	require.EqualValues(t, a1, h.Desencolar())
	require.EqualValues(t, a3, h.VerMax())
	require.EqualValues(t, a3, h.Desencolar())
	require.EqualValues(t, a2, h.VerMax())
	require.EqualValues(t, a2, h.Desencolar())
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())

}

func TestInsertarRepetidas(t *testing.T) {
	t.Log("Insertamos datos repetidos al heap")
	h := TDAheap.CrearHeap[int](func(x, y int) int {
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
		return 0
	})
	for i := 0; i < 3; i++ {
		h.Encolar(i)
		h.Encolar(9)
	}
	for i := 0; i < 3; i++ {
		require.EqualValues(t, 9, h.VerMax())
		require.EqualValues(t, 9, h.Desencolar())
	}
}
