package errores

type errorLogin struct{}

func (errorLogin) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type errorUsuarioInvalido struct{}

func (errorUsuarioInvalido) Error() string {
	return "Error: usuario no existente"
}

type errorSinUsuarioLogueado struct{}

func (errorSinUsuarioLogueado) Error() string {
	return "Error: no habia usuario loggeado"
}

type errorDarLike struct{}

func (errorDarLike) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type errorVerLike struct{}

func (errorVerLike) Error() string {
	return "Error: Post inexistente o sin likes"
}
