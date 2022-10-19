package errores

type ErrorNoEncontrado struct{}

func (e ErrorNoEncontrado) Error() string {
	return "ERROR: Lectura de archivos"
}
