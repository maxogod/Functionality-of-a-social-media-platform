package diccionario

// METODOS AUXILIARES DE ARBOL

// guardarNodo Guarda el nuevo nodo en su correspondiente lugar recursivamente (toma como garantizado que
// no tendra que actualizar).
func (a *abb[K, V]) guardarNodo(nodoActual, nuevoNodo *nodoAbb[K, V]) {
	cmp := a.cmp(nuevoNodo.clave, nodoActual.clave)
	if cmp < 0 {
		if nodoActual.izquierdo != nil {
			a.guardarNodo(nodoActual.izquierdo, nuevoNodo)
		} else {
			nodoActual.izquierdo = nuevoNodo
		}
	} else {
		if nodoActual.derecho != nil {
			a.guardarNodo(nodoActual.derecho, nuevoNodo)
		} else {
			nodoActual.derecho = nuevoNodo
		}
	}
}

// obtenerNodo Busca el nodo por clave, si la clave es mayor que la clave del nodo actual,
// busca en la derecha, sino en la izquierda. retorna nodoBuscado si lo encuentra sino nil.
func (a abb[K, V]) obtenerNodo(nodoActual, nodoPadre *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodoActual == nil {
		return nil, nil // no encontrado
	}
	cmp := a.cmp(clave, nodoActual.clave)
	if cmp > 0 {
		return a.obtenerNodo(nodoActual.derecho, nodoActual, clave) //muevo derecha
	} else if cmp < 0 {
		return a.obtenerNodo(nodoActual.izquierdo, nodoActual, clave) //muevo izquierda
	} else {
		return nodoActual, nodoPadre //encontrado
	}
}

// cantidadDeHijos devuelve la cantidad de hijos de un nodo (0 - 1 - 2)
func (n nodoAbb[K, V]) cantidadDeHijos() int {
	if n.derecho == nil && n.izquierdo == nil {
		return 0
	} else if n.derecho != nil && n.izquierdo != nil {
		return 2
	}
	return 1
}

func (a *abb[K, V]) borrarNodoSinHijos(nodoBuscado, nodoPadre *nodoAbb[K, V]) {
	if a.cantidad == 1 {
		a.raiz = nil // Caso borde raiz
	} else {
		if nodoPadre.izquierdo == nodoBuscado {
			nodoPadre.izquierdo = nil
		} else {
			nodoPadre.derecho = nil
		}
	}
}

func (a *abb[K, V]) borrarNodoUnHijo(nodoBuscado, nodoPadre *nodoAbb[K, V]) {
	var hijo *nodoAbb[K, V]
	if nodoBuscado.izquierdo != nil {
		hijo = nodoBuscado.izquierdo
	} else {
		hijo = nodoBuscado.derecho
	}

	if nodoBuscado == a.raiz { //caso borde raiz
		a.raiz = hijo
	} else if nodoPadre.izquierdo == nodoBuscado {
		nodoPadre.izquierdo = hijo
	} else {
		nodoPadre.derecho = hijo
	}
}

func (a *abb[K, V]) borrarNodoDosHijos(nodoBuscado *nodoAbb[K, V]) {
	nodoRemplazo := a.buscarReemplazo(nodoBuscado.izquierdo)

	clave := nodoRemplazo.clave
	valor := a.Borrar(clave)
	nodoBuscado.clave, nodoBuscado.valor = clave, valor
}

// buscarReemplazo a partir del nodo pasado busca el primer nodo que no tenga hijo derecho y lo devuelve
func (a abb[K, V]) buscarReemplazo(nodoRemplazo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodoRemplazo.derecho == nil {
		return nodoRemplazo
	}
	return a.buscarReemplazo(nodoRemplazo.derecho)
}

func (a abb[K, V]) iterarEntreNodos(nodoActual *nodoAbb[K, V], desde *K, hasta *K, f func(clave K, dato V) bool) bool {
	if nodoActual == nil {
		return true
	}
	return a.iterarEntreNodos(nodoActual.izquierdo, desde, hasta, f) && a.aplicarFuncionVisitar(nodoActual, desde, hasta, f) && a.iterarEntreNodos(nodoActual.derecho, desde, hasta, f)
}

// aplicarFuncionVisitar aplica la funcion visitar en caso de ser apropiado hacer esto, cuando lo hace retorna true
// si dicha funcion retorno true o false caso contrario.
func (a abb[K, V]) aplicarFuncionVisitar(nodoActual *nodoAbb[K, V], desde *K, hasta *K, f func(clave K, dato V) bool) bool {
	if desde == nil && hasta == nil {
		return f(nodoActual.clave, nodoActual.valor)
	} else if desde == nil && hasta != nil && a.cmp(*hasta, nodoActual.clave) >= 0 {
		return f(nodoActual.clave, nodoActual.valor)
	} else if hasta == nil && desde != nil && a.cmp(*desde, nodoActual.clave) <= 0 {
		return f(nodoActual.clave, nodoActual.valor)
	} else if hasta != nil && desde != nil && a.cmp(*desde, nodoActual.clave) <= 0 && a.cmp(*hasta, nodoActual.clave) >= 0 {
		return f(nodoActual.clave, nodoActual.valor)
	}
	return true
}

// METODOS AUXILIARES DE ITERADORES

// prellenarPila apila el nodo Actual junto con todos sus hijos izquierdos.
func (i *iterDic[K, V]) prellenarPila(nodoActual *nodoAbb[K, V]) {
	if nodoActual == nil {
		return
	}
	if i.desde == nil && i.hasta == nil {
		i.arbolApilado.Apilar(nodoActual)
	} else if i.desde == nil && i.cmp(*i.hasta, nodoActual.clave) >= 0 {
		i.arbolApilado.Apilar(nodoActual)
	} else if i.hasta == nil && i.cmp(*i.desde, nodoActual.clave) <= 0 {
		i.arbolApilado.Apilar(nodoActual)
	} else if i.desde != nil && i.hasta != nil && i.cmp(*i.desde, nodoActual.clave) <= 0 && i.cmp(*i.hasta, nodoActual.clave) >= 0 {
		i.arbolApilado.Apilar(nodoActual)
	} else if nodoActual.izquierdo == nil && nodoActual.derecho != nil {
		i.prellenarPila(nodoActual.derecho)
	}
	i.prellenarPila(nodoActual.izquierdo)
}
