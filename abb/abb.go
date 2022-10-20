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

// Primitivas del Arbol

func (a *abb[K, V]) Guardar(clave K, valor V) {
	nuevoNodo := &nodoAbb[K, V]{clave: clave, valor: valor}
	if a.raiz == nil {
		a.raiz = nuevoNodo // Guardar raiz
	} else if a.raiz.clave == nuevoNodo.clave {
		a.raiz.valor = nuevoNodo.valor // Actualizar raiz
	} else {
		a.guardarEntreNodos(a.raiz, nuevoNodo)
	}
}

// guardarEntreNodos Guarda el nuevo nodo en su correspondiente lugar o lo actualiza recursivamente.
func (a *abb[K, V]) guardarEntreNodos(nodoPadre, nuevoNodo *nodoAbb[K, V]) {

	if a.cmp(nuevoNodo.clave, nodoPadre.clave) < 0 {
		// Mover a Izq
		if nodoPadre.izquierdo == nil {
			// Guardar
			nuevoNodo.padre = nodoPadre
			nodoPadre.izquierdo = nuevoNodo
			a.cantidad++
		} else if nodoPadre.izquierdo.clave == nuevoNodo.clave {
			// Actualizar valor
			nodoPadre.izquierdo.valor = nuevoNodo.valor
		} else {
			a.guardarEntreNodos(nodoPadre.izquierdo, nuevoNodo)
		}
	} else if a.cmp(nuevoNodo.clave, nodoPadre.clave) > 0 {
		// Mover a Der
		if nodoPadre.derecho == nil {
			// Guardar
			nuevoNodo.padre = nodoPadre
			nodoPadre.derecho = nuevoNodo
			a.cantidad++
		} else if nodoPadre.derecho.clave == nuevoNodo.clave {
			// Actualizar valor
			nodoPadre.derecho.valor = nuevoNodo.valor
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
		// Es hijo der de su padre
		nodoBuscado.padre.derecho = nil
	} else {
		// Es hijo izq de su padre
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
		// Es hijo der de su padre
		nodoBuscado.padre.derecho = hijo
	} else {
		// Es hijo izq de su padre
		nodoBuscado.padre.izquierdo = hijo
	}
	hijo.padre = nodoBuscado.padre
}

func (a *abb[K, V]) borrarNodoDosHijos(nodoBuscado *nodoAbb[K, V]) {
	//buscamos el nodo que remplaza al nodo borrado
	nodoRemplazo := a.buscarRemplazo(nodoBuscado.izquierdo)

	clave := nodoRemplazo.clave
	valor := a.Borrar(clave)
	nodoBuscado.clave, nodoBuscado.valor = clave, valor
}

func (a abb[K, V]) buscarRemplazo(nodoRemplazo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodoRemplazo.derecho == nil {
		return nodoRemplazo
	}
	return a.buscarRemplazo(nodoRemplazo.derecho)
}

// buscarEntreNodos Busca el nodo por clave, comenzando en la raiz
// si la clave es mayor que la clave del nodo actual, busca en la derecha, sino en la izquierda.
// retorna (nodoBuscado, nil) si lo encuentra, sino (nil, errorNoEncontrado).
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

// Primitivas Iter extenos

// TODO implement

// Funcion Creacion

func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	a := new(abb[K, V])
	a.cmp = funcionCmp
	return a
}
