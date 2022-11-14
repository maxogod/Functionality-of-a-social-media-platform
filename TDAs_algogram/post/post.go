package post

type Post interface {

	//MostrarLikes muestra la cantidad de likes y quienes likearon el PostImplementation en question
	MostrarLikes() string

	//MostrarPost muestra el PostImplementation actual
	MostrarPost() string

	// ObtenerPoster Devolvemos el id del dueño del PostImplementation
	ObtenerPosterID() int

	// LikearPost le suma un like al PostImplementation, pasando por parametro la persona dando like
	LikearPost(usuarioNombre string) string

	ObtenerPostID() int
}
