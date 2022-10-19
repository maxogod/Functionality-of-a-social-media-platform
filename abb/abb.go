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
	//cmp      funcCmp[K]
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
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (a abb[K, V]) Cantidad() int {
	//TODO implement me
	panic("implement me")
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

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return new(abb[K, V])
}
