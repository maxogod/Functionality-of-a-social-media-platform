package post

import "algogram/TDAs_algogram/usuario"

type Post interface {
	//MostrarLikes muestra la cantidad de likes y quienes likearon el post en question
	MostrarLikes() string
	//MostrarPost muestra el post actual
	MostrarPost() string
	//Devolvemos el dueño del post
	DevolverPosteador() usuario.Usuario
	LikearPost(usuarioNombre string) string
}
