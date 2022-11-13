package post

type Post interface {

	//MostrarLikes muestra la cantidad de likes y quienes likearon el post en question
	MostrarLikes() string

	//MostrarPost muestra el post actual
	MostrarPost() string

	// ObtenerPoster Devolvemos el id del dueño del post
	ObtenerPosterID() int

	// LikearPost le suma un like al post, pasando por parametro la persona dando like
	LikearPost(usuarioNombre string) string
}
