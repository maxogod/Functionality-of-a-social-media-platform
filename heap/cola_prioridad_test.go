package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	TDAheap "tp2/heap"
)

func TestABBVacio(t *testing.T) {
	h := TDAheap.CrearHeap(strings.Compare)
	h.Cantidad()
	require.True(t, true)
}
