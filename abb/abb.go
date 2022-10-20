package diccionario

import (
	"diccionario/errores"
	dic "diccionario/hash"
)

type nodoAbb[K comparable, V any] struct {
	padre     *nodoAbb[K, V]
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
	a.cantidad++
}

// guardarEntreNodos Guarda el nuevo nodo en su correspondiente lugar usando recursividad.
func (a *abb[K, V]) guardarEntreNodos(nodoPadre, nuevoNodo *nodoAbb[K, V]) {
	if a.raiz == nil {
		a.raiz = nuevoNodo
	} else if a.raiz.clave == nuevoNodo.clave {
		a.raiz.valor = nuevoNodo.valor // actualizar raiz
	} else if a.cmp(nuevoNodo.clave, nodoPadre.clave) < 0 {
		// Mover a Izq
		if nodoPadre.izquierdo == nil {
			nuevoNodo.padre = nodoPadre
			nodoPadre.izquierdo = nuevoNodo
		} else if nodoPadre.izquierdo.clave == nuevoNodo.clave {
			nodoPadre.izquierdo.valor = nuevoNodo.valor // actualizar valor
		} else {
			a.guardarEntreNodos(nodoPadre.izquierdo, nuevoNodo)
		}
	} else if a.cmp(nuevoNodo.clave, nodoPadre.clave) > 0 {
		// Mover a Der
		if nodoPadre.derecho == nil || nodoPadre.derecho.clave == nuevoNodo.clave {
			nuevoNodo.padre = nodoPadre
			nodoPadre.derecho = nuevoNodo
		} else if nodoPadre.derecho.clave == nuevoNodo.clave {
			nodoPadre.derecho.valor = nuevoNodo.valor // actualizar valor
		} else {
			a.guardarEntreNodos(nodoPadre.derecho, nuevoNodo)
		}
	}
}

func (a abb[K, V]) Pertenece(clave K) bool {
	_, err := a.buscarEntreNodos(a.raiz, clave)
	if err != nil {
		return false
	}
	return true
}

func (a abb[K, V]) Obtener(clave K) V {
	nodoBuscado, err := a.buscarEntreNodos(a.raiz, clave)
	if err != nil {
		panic(err.Error())
	}
	return nodoBuscado.valor
}

//Comparamos cada uno de los nodos del arbol, comenzando en la raiz
//si la clave es mayor que la clave del nodo actual, nos movemos a la derecha,
//sino, nos movemos a la izquierda
//en caso de no cumplir nada de lo previamente mencionado, devuelve error
//tambien devuelve error cuando el arbol está vacio
func (a abb[K, V]) buscarEntreNodos(nodoPadre *nodoAbb[K, V], clave K) (*nodoAbb[K, V], error) {
	if nodoPadre == nil {
		return nil, new(errores.ErrorNoEncontrado)
	} else if a.cmp(clave, nodoPadre.clave) > 0 {
		//muevo Der
		return a.buscarEntreNodos(nodoPadre.derecho, clave)
	} else if a.cmp(clave, nodoPadre.clave) < 0 {
		//muevo Izq
		return a.buscarEntreNodos(nodoPadre.izquierdo, clave)
	}
	return nodoPadre, nil
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodoBuscado, err := a.buscarEntreNodos(a.raiz, clave)
	if err != nil {
		panic(err.Error())
	}

	// Nodo existe y sera borrado
	valor := nodoBuscado.valor
	a.cantidad--
	if nodoBuscado.izquierdo == nil && nodoBuscado.derecho == nil {
		// Sin hijos
		a.borrarNodoSinHijos(nodoBuscado)
	} else if nodoBuscado.izquierdo != nil && nodoBuscado.derecho != nil {
		// 2 hijos
		a.borrarNodoDosHijos(nodoBuscado)
	} else {
		// 1 hijo
		a.borrarNodoUnHijo(nodoBuscado)
	}
	return valor
}

func (a *abb[K, V]) borrarNodoSinHijos(nodoBuscado *nodoAbb[K, V]) {
	if nodoBuscado.padre == nil {
		// Caso borde raiz
		a.raiz = nil
	} else if a.cmp(nodoBuscado.clave, nodoBuscado.padre.clave) > 0 {
		nodoBuscado.padre.derecho = nil
	} else {
		nodoBuscado.padre.izquierdo = nil
	}
}

func (a *abb[K, V]) borrarNodoUnHijo(nodoBuscado *nodoAbb[K, V]) {
	var hijo *nodoAbb[K, V]
	if nodoBuscado.izquierdo != nil {
		hijo = nodoBuscado.izquierdo
	} else {
		hijo = nodoBuscado.derecho
	}
	if nodoBuscado.padre == nil {
		// Caso borde raiz
		a.raiz = hijo
		hijo.padre = nil
	} else if a.cmp(nodoBuscado.clave, nodoBuscado.padre.clave) > 0 {
		nodoBuscado.padre.derecho = hijo
	} else {
		nodoBuscado.padre.izquierdo = hijo
	}
	hijo.padre = nodoBuscado.padre
}

func (a *abb[K, V]) borrarNodoDosHijos(nodoBuscado *nodoAbb[K, V]) {
	//buscamos el nodo que remplaza al nodo borrado
	nodoRemplazo := a.buscarRemplazo(nodoBuscado.izquierdo)
	if nodoBuscado.padre == nil {
		// Caso borde raiz
		nodoRemplazo.derecho = a.raiz.derecho
		nodoRemplazo.padre = nil
		a.raiz = nodoRemplazo
	} else {
		nodoRemplazo.padre.derecho = nodoRemplazo.izquierdo
	}
	if nodoRemplazo.izquierdo != nil {
		//antes de remplazar el nodo, en caso que el remplazo tenga hijo izq
		nodoRemplazo.izquierdo.padre = nodoRemplazo.padre
	}
	nodoBuscado.clave, nodoBuscado.valor = nodoRemplazo.clave, nodoRemplazo.valor
}

func (a abb[K, V]) buscarRemplazo(nodoRemplazo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodoRemplazo.derecho == nil {
		return nodoRemplazo
	}
	return a.buscarRemplazo(nodoRemplazo.derecho)
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
