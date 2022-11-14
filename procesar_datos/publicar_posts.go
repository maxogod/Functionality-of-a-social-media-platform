package procesar_datos

import (
	diccionario "algogram/TDAs/hash"
	"algogram/TDAs_algogram/post"
	"algogram/TDAs_algogram/usuario"
)

func GuardarPostEnFeeds(nuevoPost post.Post, usuarios diccionario.Diccionario[string, usuario.Usuario]) {
	usuarios.Iterar(
		func(K string, V usuario.Usuario) bool {
			V.AgregarPostFeed(nuevoPost)
			return true
		})
}
