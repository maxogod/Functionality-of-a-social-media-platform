package diccionario

import dic "diccionario/TP2/hash"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func (a abb[K, V]) Guardar(clave K, dato V) {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Pertenece(clave K) bool {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Obtener(clave K) V {
	//Si el arbol esta vacio, tira error
	if a.Cantidad() == 0 {
		error()
	} else if a.cmp(a.raiz.clave, clave) == 0 { //si el dato esta en la raiz, devolvemos la raiz
		return a.raiz.dato
	} else if a.cmp(clave, a.raiz.clave) > 0 && a.raiz.derecho != nil { //si la clave es mayor que el nodo raiz, nos movemos a la derecha del arbol y tambien debe EXISTIR un nodo a la derecha
		//muevo Der
		return a.raiz.derecho.obtenerEntreNodos(a.cmp, clave)

	} else if a.cmp(clave, a.raiz.clave) < 0 && a.raiz.izquierdo != nil { //si la clave es menor que el nodo raiz, nos movemos a la izquierda del arbol y tambien debe EXISTIR un nodo a la izquierda
		//muevo Izq
		return a.raiz.izquierdo.obtenerEntreNodos(a.cmp, clave)

	}
	//en caso de no cumplir nada de lo previamente mencionado, tira error
	error()
	return nil
}

func (n nodoAbb[K, V]) obtenerEntreNodos(cmp func(K, K) int, clave K) V {
	if cmp(n.clave, clave) == 0 { //si la clave del nodo actual es igual a la clave en cuestion, nos devuelve el dato del nodo
		return n.dato
	} else if cmp(clave, n.clave) > 0 && n.derecho != nil { //comparamos y vemos de movernos al nodo a la derecha (en caso que exista un nodo a la derecha)
		return n.derecho.obtenerEntreNodos(cmp, clave)

	} else if cmp(clave, n.clave) < 0 && n.izquierdo != nil { //comparamos y vemos de movernos al nodo a la izquierda (en caso que exista un nodo a la izquierda)
		//muevo Izq
		return n.izquierdo.obtenerEntreNodos(cmp, clave)

	}
	//en caso de no cumplir nada de lo previamente mencionado, tira error
	error()
	return nil
}

func error() {
	panic("dato no esta")
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
