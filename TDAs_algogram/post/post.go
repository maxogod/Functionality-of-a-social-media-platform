package post

import "algogram/TDAs_algogram/usuario"

type Post interface {

	//MostrarLikes muestra la cantidad de likes y quienes likearon el post en question
	MostrarLikes() string

	//MostrarPost muestra el post actual
	MostrarPost() string

	// ObtenerPoster Devolvemos el dueño del post
	ObtenerPoster() usuario.Usuario

	// LikearPost le suma un like al post, pasando por parametro la persona dando like
	LikearPost(usuarioNombre string) string
}
