package cola_prioridad

import "cola_prioridad/errores"

const (
	_LARGO_MINIMO       = 20
	_CUADRIPLICADOR     = 4
	_FACTOR_REDIMENSION = 2
)

type heap[T comparable] struct {
	cantidad int
	datos    []T
	cmp      func(T, T) int
}

// Funciones de creacion

func CrearHeap[T comparable](funcionCmp func(T, T) int) ColaPrioridad[T] {
	h := new(heap[T])
	h.datos = make([]T, _LARGO_MINIMO)
	h.cmp = funcionCmp
	return h
}

func CrearHeapArr[T comparable](arreglo []T, funcionCmp func(T, T) int) ColaPrioridad[T] {
	h := new(heap[T])
	tamanio := len(arreglo) * 2
	if tamanio == 0 {
		tamanio = _LARGO_MINIMO
	}
	h.datos = make([]T, tamanio)
	copy(h.datos, arreglo)
	h.cantidad = len(arreglo)
	h.cmp = funcionCmp
	heapify(h.datos, h.cmp, h.cantidad)
	return h
}

// Implementacion de heap

func (h heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(valor T) {
	h.datos[h.cantidad] = valor
	h.cantidad++
	downHeap(h.datos, h.cmp, h.cantidad, h.cantidad-1)
	if h.cantidad == len(h.datos) {
		h.redimencionar(len(h.datos) * _FACTOR_REDIMENSION)
	}
}

func (h heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic(errores.ErrorColaVacia{}.Error())
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic(errores.ErrorColaVacia{}.Error())
	}
	dato := h.datos[0]
	h.datos[0] = h.datos[h.cantidad-1]
	h.cantidad--
	downHeap(h.datos, h.cmp, h.cantidad, 0)
	if h.cantidad*_CUADRIPLICADOR <= len(h.datos) && h.cantidad*_CUADRIPLICADOR >= _LARGO_MINIMO {
		h.redimencionar(len(h.datos) / _FACTOR_REDIMENSION)
	}
	return dato
}

func (h heap[T]) Cantidad() int {
	return h.cantidad
}

func (h *heap[T]) redimencionar(nuevoLargo int) {
	newArr := make([]T, nuevoLargo)
	copy(newArr, h.datos)
	h.datos = newArr
}

// Funciones adicionales

func HeapSort[T comparable](elementos []T, funcionCmp func(T, T) int) {
	heapify(elementos, funcionCmp, len(elementos))
	for cant := len(elementos) - 1; cant > 0; cant-- {
		elementos[0], elementos[cant] = elementos[cant], elementos[0]
		downHeap(elementos, funcionCmp, cant, 0)
	}
}

func downHeap[T comparable](datos []T, cmp func(T, T) int, cantidad, posAEvaluar int) {
	if posAEvaluar < 0 {
		return
	}
	izq := 2*posAEvaluar + 1
	der := 2*posAEvaluar + 2
	maxPos := max[T](datos, cmp, cantidad, posAEvaluar, der, izq)

	if maxPos != posAEvaluar {
		datos[maxPos], datos[posAEvaluar] = datos[posAEvaluar], datos[maxPos]
		downHeap(datos, cmp, cantidad, maxPos)
	}
	downHeap(datos, cmp, cantidad, maxPos-1)
}

func heapify[T comparable](datos []T, cmp func(T, T) int, cantidad int) {
	for i := cantidad - 1; i > -1; i-- {
		downHeap(datos, cmp, cantidad, i)
	}
}

func max[T comparable](datos []T, cmp func(T, T) int, cantidad, indexPadre, indexIzq, indexDer int) int {
	var max = indexPadre
	if indexIzq < cantidad && cmp(datos[indexIzq], datos[max]) > 0 {
		max = indexIzq
	}
	if indexDer < cantidad && cmp(datos[indexDer], datos[max]) > 0 {
		max = indexDer
	}
	return max
}
