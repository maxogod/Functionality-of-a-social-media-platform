package erroresAbb

type ErrorNoEncontrado struct{}

func (e ErrorNoEncontrado) Error() string {
	return "La clave no pertenece al diccionario"
}

type ErrorIterTermino struct{}

func (e ErrorIterTermino) Error() string {
	return "El iterador termino de iterar"
}
