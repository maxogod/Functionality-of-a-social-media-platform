package lista

type nodo[T any] struct {
	dato     T
	proximo  *nodo[T]
	anterior *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	lista          *listaEnlazada[T]
	posicionActual *nodo[T]
}

// Metodos de listaEnlazada

func (l *listaEnlazada[T]) errores() {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) crearNodo(nuevoDato T) *nodo[T] {
	return &nodo[T]{dato: nuevoDato}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		prox := l.primero
		l.primero = nuevoNodo
		l.primero.proximo = prox
		l.primero.proximo.anterior = l.primero
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(nuevoDato T) {
	nuevoNodo := l.crearNodo(nuevoDato)

	if l.EstaVacia() {
		l.primero = nuevoNodo
		l.ultimo = nuevoNodo
	} else {
		ant := l.ultimo
		l.ultimo = nuevoNodo
		l.ultimo.anterior = ant
		l.ultimo.anterior.proximo = l.ultimo
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	l.errores()
	const _LARGO_MIN = 1
	dato := l.primero.dato

	if l.largo == _LARGO_MIN {
		l.primero = nil
		l.ultimo = nil
	} else {
		l.primero = l.primero.proximo
		l.primero.anterior = l.primero.anterior.anterior
	}
	l.largo--

	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	l.errores()
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	l.errores()
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

// Iterar - iterador Interno
func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for l.primero != nil && visitar(l.primero.dato) {
		l.primero = l.primero.proximo
	}
}

// Iterador - Crea un iterador externo
func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iteradorListaEnlazada[T])
	iter.lista = l
	iter.posicionActual = l.primero
	return iter
}

//Metodos de iteradorListaEnlazada (EXTERNO)

func (i *iteradorListaEnlazada[T]) errorDeIterador() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (i *iteradorListaEnlazada[T]) VerActual() T {
	i.errorDeIterador()
	return i.posicionActual.dato
}

func (i *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return i.posicionActual != nil
}

func (i *iteradorListaEnlazada[T]) Siguiente() T {
	i.errorDeIterador()
	actual := i.posicionActual.dato
	i.posicionActual = i.posicionActual.proximo
	return actual
}

func (i *iteradorListaEnlazada[T]) Insertar(nuevoDato T) {

	if i.lista.EstaVacia() {
		// Vacia
		i.lista.InsertarPrimero(nuevoDato)
		i.posicionActual = i.lista.primero
	} else if i.posicionActual == nil {
		// Final
		i.lista.InsertarUltimo(nuevoDato)
		i.posicionActual = i.lista.ultimo
	} else if i.posicionActual.anterior == nil {
		// Principio
		i.lista.InsertarPrimero(nuevoDato)
		i.posicionActual = i.lista.primero
	} else {
		// Medio
		nuevoNodo := i.lista.crearNodo(nuevoDato)

		prox := i.posicionActual
		i.posicionActual = nuevoNodo
		i.posicionActual.proximo = prox
		i.posicionActual.anterior = prox.anterior
		i.posicionActual.proximo.anterior = i.posicionActual
		i.posicionActual.anterior.proximo = i.posicionActual
		i.lista.largo++
	}
}

func (i *iteradorListaEnlazada[T]) Borrar() T {
	i.errorDeIterador()
	i.lista.errores()
	dato := i.posicionActual.dato

	if i.posicionActual.anterior == nil {
		// Principio
		i.lista.BorrarPrimero()
		i.Siguiente()
	} else if i.posicionActual.proximo == nil {
		// Final
		i.posicionActual.anterior.proximo = nil
		i.lista.ultimo = i.lista.ultimo.anterior
		i.posicionActual = nil
		i.lista.largo--
	} else {
		// Medio
		i.posicionActual.anterior.proximo = i.posicionActual.proximo
		i.posicionActual.proximo.anterior = i.posicionActual.anterior
		i.posicionActual = i.posicionActual.proximo
		i.lista.largo--
	}
	return dato
}

// CrearListaEnlazada - Funcion creadora de lista
func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}
