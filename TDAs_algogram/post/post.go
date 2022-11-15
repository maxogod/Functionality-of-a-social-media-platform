package post

type Post interface {

	//MostrarLikes muestra la cantidad de likes y quienes likearon el postImplementation en question
	MostrarLikes() (string, error)

	//MostrarPost muestra el postImplementation actual
	MostrarPost() string

	// ObtenerPoster Devolvemos el id del dueño del postImplementation
	ObtenerPosterID() int

	// LikearPost le suma un like al postImplementation, pasando por parametro la persona dando like
	LikearPost(usuarioNombre string) string

	ObtenerPostID() int
}
