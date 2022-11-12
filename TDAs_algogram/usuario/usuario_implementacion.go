package usuario

import (
	hp "algogram/TDAs/heap"
	post "algogram/TDAs_algogram/post"
	"algogram/errores"
	"strings"
)

type usuario struct {
	feed   hp.ColaPrioridad[post.Post]
	nombre string
}

// CrearUsuario Funcion de creacion de un usuario
func CrearUsuario(nombre string) Usuario {
	usuarioCreado := new(usuario)
	usuarioCreado.feed = hp.CrearHeap[post.Post](cmp)
	usuarioCreado.nombre = nombre
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

// Funciones adicionales abajo

// cmp Funcion de comparacion basada en afinidad de usuarios
func cmp(post1, post2 post.Post) int {
	if strings.Compare(post1.ObtenerPoster().nombre, post2.ObtenerPoster().nombre) > 1 {
		return 1
	} else if strings.Compare(post1.ObtenerPoster().nombre, post2.ObtenerPoster().nombre) < 1 {
		return -1
	}
	return 0
}
