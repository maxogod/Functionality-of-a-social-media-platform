package usuario

import (
	hp "algogram/TDAs/heap"
	"algogram/TDAs_algogram/post"
	"algogram/errores"
	"math"
)

type usuario struct {
	feed   hp.ColaPrioridad[postFeed]
	nombre string
	id     int
}

type postFeed struct {
	idPost      int
	idPosteador int
}

// CrearUsuario Funcion de creacion de un usuario
func CrearUsuario(nombre string, id int) Usuario {
	usuarioCreado := new(usuario)
	usuarioCreado.feed = hp.CrearHeap[postFeed](usuarioCreado.afinidad)
	usuarioCreado.nombre = nombre
	usuarioCreado.id = id
	return usuarioCreado
}

// Metodos de usuario abajo

func (u *usuario) VerSigPost() (int, error) {
	if u.feed.EstaVacia() {
		return -1, errores.ErrorVerPost{}
	}
	return u.feed.Desencolar().idPost, nil
}

func (u *usuario) AgregarPostFeed(post post.Post) {
	pF := new(postFeed)
	pF.idPosteador = post.ObtenerPosterID()
	pF.idPost = post.ObtenerPostID()

	u.feed.Encolar(*pF)
}

func (u *usuario) ObtenerNombre() string {
	return u.nombre
}

func (u *usuario) ObtenerId() int {
	return u.id
}

// afinidad es una funcion de comparacion basada en afinidad de usuarios
func (u *usuario) afinidad(post1, post2 postFeed) int {
	x := int(math.Abs(float64(u.id) - float64(post1.idPosteador)))
	y := int(math.Abs(float64(u.id) - float64(post2.idPosteador)))
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}
