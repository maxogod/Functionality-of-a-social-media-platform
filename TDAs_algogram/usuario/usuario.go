package usuario

import "algogram/TDAs_algogram/post"

type Usuario interface {

	// VerSigPost devuelve el siguiente post disponible para ver de un usuario basado en la afinidad que este tenga
	// con el poster
	VerSigPost() (string, error)

	// AgregarPostFeed agrega el post pasado por parametro a la feed de el usuario
	AgregarPostFeed(post post.Post)

	// ObtenerNombre devuelve el nombre del usuario
	ObtenerNombre() string

	// ObtenerId devuelve el id del usuario
	ObtenerId() int
}
