package cola_prioridad

type heap[T comparable] struct {
	cantidad int
	datos    []T
	cmp      func(T, T) int
}

// Funciones de creacion

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	c := new(heap[T])
	return c
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	panic("Implement me")
}

// Implementacion de heap

func (h heap[T]) EstaVacia() bool {
	//TODO implement me
	panic("implement me")
}

func (h heap[T]) Encolar(t T) {
	//TODO implement me
	panic("implement me")
}

func (h heap[T]) VerMax() T {
	//TODO implement me
	panic("implement me")
}

func (h heap[T]) Desencolar() T {
	//TODO implement me
	panic("implement me")
}

func (h heap[T]) Cantidad() int {
	return h.cantidad
}

// Funciones adicionales

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	panic("Implement me")
}
