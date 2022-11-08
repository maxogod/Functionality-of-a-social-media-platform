package erroresHeap

type ErrorColaVacia struct{}

func (e ErrorColaVacia) Error() string {
	return "La cola esta vacia"
}
