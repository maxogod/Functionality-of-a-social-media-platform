package usuario

import (
	hp "algogram/TDAs/heap"
	"algogram/TDAs_algogram/post"
	"algogram/errores"
	"math"
)

type usuario struct {
	feed   hp.ColaPrioridad[post.Post]
	nombre string
	id     int
}

// CrearUsuario Funcion de creacion de un usuario
func CrearUsuario(nombre string, id int) Usuario {
	usuarioCreado := new(usuario)
	usuarioCreado.feed = hp.CrearHeap[post.Post](usuarioCreado.afinidad)
	usuarioCreado.nombre = nombre
	usuarioCreado.id = id
	return usuarioCreado
}

// Metodos de usuario abajo

func (u *usuario) VerSigPost() (string, error) {
	if u.feed.EstaVacia() {
		return "", errores.ErrorVerPost{}
	}
	return u.feed.Desencolar().MostrarPost(), nil
}

func (u *usuario) AgregarPostFeed(post post.Post) {
	u.feed.Encolar(post)
}

func (u *usuario) ObtenerNombre() string {
	return u.nombre
}

func (u *usuario) ObtenerId() int {
	return u.id
}

// Funciones adicionales abajo

// afinidad es una funcion de comparacion basada en afinidad de usuarios
func (u *usuario) afinidad(post1, post2 post.Post) int {
	x := int(math.Abs(float64(u.id) - float64(post1.ObtenerPoster().ObtenerId())))
	y := int(math.Abs(float64(u.id) - float64(post2.ObtenerPoster().ObtenerId())))
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}
