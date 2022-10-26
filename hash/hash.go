package diccionario

import (
	"fmt"
	"tp2/lista"
)

const (
	_LONGITUD_INICIAL     = 23
	_REDIMENSION_AGRANDAR = 5
	_CUADRUPLE            = 4
)

type hashDato[K comparable, V any] struct {
	clave K
	valor V
}

type hashMap[K comparable, V any] struct {
	hashArray []lista.Lista[hashDato[K, V]]
	longitud  int
}

type iteradorHash[K comparable, V any] struct {
	hashArray    []lista.Lista[hashDato[K, V]]
	index        int
	subListaIter lista.IteradorLista[hashDato[K, V]]
}

// Funcion de creacion de hash

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	h := new(hashMap[K, V])
	h.hashArray = make([]lista.Lista[hashDato[K, V]], _LONGITUD_INICIAL)
	for i := range h.hashArray {
		h.hashArray[i] = lista.CrearListaEnlazada[hashDato[K, V]]()
	}
	return h
}

// Implementacion de HashMap

func (h *hashMap[K, V]) Guardar(clave K, valor V) {
	if h.Pertenece(clave) {
		h.actualizar(clave, valor)
		return
	}
	nuevoDato := &hashDato[K, V]{clave: clave, valor: valor}
	index := convertir(clave, len(h.hashArray))
	h.hashArray[index].InsertarPrimero(*nuevoDato)
	h.longitud++
	if h.hashArray[index].Largo() >= _REDIMENSION_AGRANDAR {
		h.redimesionar(proxPrimo(len(h.hashArray) * 2))
	}
}

func (h *hashMap[K, V]) actualizar(clave K, valorActualizado V) {
	index := convertir(clave, len(h.hashArray))
	listaIndex := h.hashArray[index]
	for iter := listaIndex.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			iter.Borrar()
			iter.Insertar(hashDato[K, V]{clave: clave, valor: valorActualizado})
		}
	}
}

func (h hashMap[K, V]) Pertenece(clave K) bool {
	index := convertir(clave, len(h.hashArray))
	listaIndex := h.hashArray[index]
	for iter := listaIndex.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			return true
		}
	}
	return false
}

func (h *hashMap[K, V]) Obtener(clave K) V {
	index := convertir(clave, len(h.hashArray))
	subLista := h.hashArray[index]
	for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			return iter.VerActual().valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashMap[K, V]) Borrar(clave K) V {
	index := convertir(clave, len(h.hashArray))
	subLista := h.hashArray[index]
	for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			dato := iter.Borrar()
			h.longitud--
			if h.longitud*_CUADRUPLE <= len(h.hashArray) && h.longitud*_CUADRUPLE > _LONGITUD_INICIAL {
				h.redimesionar(proxPrimo(h.longitud / 2))
			}
			return dato.valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h hashMap[K, V]) Cantidad() int {
	return h.longitud
}

func (h hashMap[K, V]) Iterar(f func(clave K, valor V) bool) {
	for _, subLista := range h.hashArray {
		for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			dato := iter.VerActual()
			if !f(dato.clave, dato.valor) {
				return
			}
		}
	}
}

func (h hashMap[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorHash[K, V])
	iter.hashArray = h.hashArray
	for iter.index < len(iter.hashArray) && iter.hashArray[iter.index].EstaVacia() {
		iter.index++
	}
	if iter.index == len(iter.hashArray) {
		iter.index = 0
	}
	iter.subListaIter = iter.hashArray[iter.index].Iterador()
	return iter
}

func (h *hashMap[K, V]) redimesionar(nuevoLen int) {
	nuevoHash := new(hashMap[K, V])
	nuevoHash.hashArray = make([]lista.Lista[hashDato[K, V]], nuevoLen)
	for i := range nuevoHash.hashArray {
		nuevoHash.hashArray[i] = lista.CrearListaEnlazada[hashDato[K, V]]()
	}
	for _, subLista := range h.hashArray {
		for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			nuevoHash.Guardar(iter.VerActual().clave, iter.VerActual().valor)
		}
	}
	h.hashArray = nuevoHash.hashArray
}

// Implementacion de iter Externo

func (i iteradorHash[K, V]) HaySiguiente() bool {

	if i.subListaIter == nil {
		return false
	}
	if i.subListaIter.HaySiguiente() {
		return true
	}

	for i.hashArray[i.index].EstaVacia() {
		i.index++
		if i.index == len(i.hashArray) {
			return false
		}
	}
	return true
}

func (i iteradorHash[K, V]) VerActual() (K, V) {
	if i.subListaIter != nil {
		return i.subListaIter.VerActual().clave, i.subListaIter.VerActual().valor
	}
	panic("El iterador termino de iterar")
}

func (i *iteradorHash[K, V]) Siguiente() K {

	if i.subListaIter != nil && i.subListaIter.HaySiguiente() {
		clave := i.subListaIter.Siguiente().clave
		if !i.subListaIter.HaySiguiente() {
			i.proxIndexOcupado()
		}
		return clave
	} else if i.HaySiguiente() {
		i.proxIndexOcupado()
		return i.subListaIter.VerActual().clave
	}
	panic("El iterador termino de iterar")
}

// funcion de hashing + Otras funciones privadas

func sdbmHash(data []byte, longitud int) int {
	// documentacion: https://www.programmingalgorithms.com/algorithm/sdbm-hash/c/
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}
	return int(hash) % longitud
}

func (i *iteradorHash[K, V]) proxIndexOcupado() {
	i.index++
	for i.index < len(i.hashArray) && i.hashArray[i.index].EstaVacia() {
		i.index++
	}
	if i.index == len(i.hashArray) {
		i.subListaIter = nil
	} else {
		i.subListaIter = i.hashArray[i.index].Iterador()
	}
}

func convertir(T any, longitud int) int {
	dato := convertirABytes(T)
	index := sdbmHash(dato, longitud)
	if index < 0 {
		index *= -1
	}
	return index
}

func convertirABytes(clave any) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func proxPrimo(n int) int {
	if n <= 1 {
		return 2
	}
	primo := n
	encontrado := false
	for !encontrado {
		primo += 1
		if esPrimo(primo) {
			encontrado = true
		}
	}
	return primo
}
