package errores

type ErrorNoEncontrado struct{}

func (e ErrorNoEncontrado) Error() string {
	return "La clave no pertenece al diccionario"
}
