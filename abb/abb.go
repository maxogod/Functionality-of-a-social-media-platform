package diccionario

import (
	"diccionario/TP2/errores"
	dic "diccionario/TP2/hash"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	valor     V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func (a *abb[K, V]) Guardar(clave K, valor V) {
	nuevoNodo := &nodoAbb[K, V]{clave: clave, valor: valor}
	a.guardarEntreNodos(a.raiz, nuevoNodo)
}

// guardarEntreNodos Guarda el nuevo nodo en su correspondiente lugar usando recursividad.
func (a *abb[K, V]) guardarEntreNodos(nodoPadre, nuevoNodo *nodoAbb[K, V]) {
	if nodoPadre == nil || nodoPadre.clave == nuevoNodo.clave {
		nodoPadre = nuevoNodo // Si esta vacio O si hay que actualizarlo
	} else if a.cmp(nuevoNodo.clave, nodoPadre.clave) < 0 {
		// Mover a Izq
		a.guardarEntreNodos(nodoPadre.izquierdo, nuevoNodo)
	} else if a.cmp(nuevoNodo.clave, nodoPadre.clave) > 0 {
		// Mover a Der
		a.guardarEntreNodos(nodoPadre.derecho, nuevoNodo)
	}
}

func (a abb[K, V]) Pertenece(clave K) bool {
	_, err := a.obtenerEntreNodos(a.raiz, clave)
	if err != nil {
		return false
	}
	return true
}

func (a abb[K, V]) Obtener(clave K) V {

	valor, _ := a.obtenerEntreNodos(a.raiz, clave)
	return valor
}

/*
Comparamos cada uno de los nodos del arbol, comenzando en la raiz obviamente
si la clave es mayor que la clave del nodo actual y tambien que EXISTA un nodo derecho, nos movemos a la derecha,
si es menor y existe un nodo a la izquierda nos movemos a la izquierda
en caso de no cumplir nada de lo previamente mencionado, tira error
tambien tira error cuando el arbol está vacia
*/
func (a abb[K, V]) obtenerEntreNodos(nodoPadre *nodoAbb[K, V], clave K) (V, error) {
	if nodoPadre == nil {
		return nil, new(errores.ErrorNoEncontrado)
	} else if a.cmp(nodoPadre.clave, clave) == 0 {
		return nodoPadre.valor, nil
	} else if a.cmp(clave, nodoPadre.clave) > 0 {
		//muevo Der
		return a.obtenerEntreNodos(nodoPadre.derecho, clave)
	} else if a.cmp(clave, nodoPadre.clave) < 0 {
		//muevo Izq
		a.raiz = a.raiz.izquierdo
		return a.obtenerEntreNodos(nodoPadre.izquierdo, clave)
	}
	//en caso de no cumplir nada de lo previamente mencionado, tira error
	return nil, new(errores.ErrorNoEncontrado)
}

func (a abb[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Iterador() dic.IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) IteradorRango(desde *K, hasta *K) dic.IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	a := new(abb[K, V])
	a.cmp = funcionCmp
	return a
}
