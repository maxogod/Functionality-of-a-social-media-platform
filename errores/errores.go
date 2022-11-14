package errores

type ErrorLogin struct{}

func (ErrorLogin) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInvalido struct{}

func (ErrorUsuarioInvalido) Error() string {
	return "Error: usuario no existente"
}

type ErrorSinUsuarioLogueado struct{}

func (ErrorSinUsuarioLogueado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorDarLike struct{}

func (ErrorDarLike) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorVerLike struct{}

func (ErrorVerLike) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErrorVerPost struct{}

func (ErrorVerPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorLecturaArchivos struct{}

func (ErrorLecturaArchivos) Error() string {
	return "Error: Lectura de archivos"
}

type ErrorParametros struct{}

func (ErrorParametros) Error() string {
	return "Error: Faltan parámetros"
}
