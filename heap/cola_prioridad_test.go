package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	TDAheap "tp2/heap"
	"tp2/heap/errores"
)

func TestHeapVacio(t *testing.T) {
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
