package diccionario

import (
	errores "algogram/TDAs/abb/erroresAbb"
	dic "algogram/TDAs/hash"
	"algogram/TDAs/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	valor     V
}

type iterDic[K comparable, V any] struct {
	arbolApilado pila.Pila[*nodoAbb[K, V]]
	desde        *K
	hasta        *K
	cmp          func(K, K) int
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

// Funcion creacion abb

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	a := new(abb[K, V])
	a.cmp = funcionCmp
	return a
}

// Primitivas del Arbol

func (a *abb[K, V]) Guardar(clave K, valor V) {
	nuevoNodo := &nodoAbb[K, V]{clave: clave, valor: valor}
	nodoYaExiste, _ := a.obtenerNodo(a.raiz, nil, clave) // nil si no existe

	if nodoYaExiste != nil {
		nodoYaExiste.valor = valor // Actualizar nodo
	} else if a.raiz == nil {
		a.raiz = nuevoNodo // Guardar raiz
		a.cantidad++
	} else {
		a.guardarNodo(a.raiz, nuevoNodo) // Guardar nodo
		a.cantidad++
	}
}

func (a abb[K, V]) Pertenece(clave K) bool {
	nodoBuscado, _ := a.obtenerNodo(a.raiz, nil, clave)
	return nodoBuscado != nil
}

func (a abb[K, V]) Obtener(clave K) V {
	nodoBuscado, _ := a.obtenerNodo(a.raiz, nil, clave)
	if nodoBuscado == nil {
		panic(new(errores.ErrorNoEncontrado).Error())
	}
	return nodoBuscado.valor
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodoBuscado, nodoPadre := a.obtenerNodo(a.raiz, nil, clave)
	if nodoBuscado == nil {
		panic(new(errores.ErrorNoEncontrado).Error())
	}
	valor := nodoBuscado.valor
	cantHijos := nodoBuscado.cantidadDeHijos()
	switch cantHijos {
	case 0:
		a.borrarNodoSinHijos(nodoBuscado, nodoPadre)
		a.cantidad--
	case 1:
		a.borrarNodoUnHijo(nodoBuscado, nodoPadre)
		a.cantidad--
	case 2:
		a.borrarNodoDosHijos(nodoBuscado)
	}
	return valor
}

func (a abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	a.iterarEntreNodos(a.raiz, nil, nil, f)
}

func (a abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarEntreNodos(a.raiz, desde, hasta, visitar)
}

func (a abb[K, V]) Iterador() dic.IterDiccionario[K, V] {
	i := new(iterDic[K, V])
	i.arbolApilado = pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	i.prellenarPila(a.raiz)
	return i
}

func (a abb[K, V]) IteradorRango(desde *K, hasta *K) dic.IterDiccionario[K, V] {
	i := new(iterDic[K, V])
	i.cmp = a.cmp
	i.desde = desde
	i.hasta = hasta
	i.arbolApilado = pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	i.prellenarPila(a.raiz)
	return i
}

// Primitivas Iter externos

func (i iterDic[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic(new(errores.ErrorIterTermino).Error())
	}
	return i.arbolApilado.VerTope().clave, i.arbolApilado.VerTope().valor
}

func (i iterDic[K, V]) Siguiente() K {
	if !i.HaySiguiente() {
		panic(new(errores.ErrorIterTermino).Error())
	}
	elem := i.arbolApilado.Desapilar()
	i.prellenarPila(elem.derecho)
	return elem.clave
}

func (i iterDic[K, V]) HaySiguiente() bool {
	return !i.arbolApilado.EstaVacia()
}
