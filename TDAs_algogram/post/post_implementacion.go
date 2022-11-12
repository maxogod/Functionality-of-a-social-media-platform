package post

import (
	TDAdic "algogram/TDAs/abb"
	"algogram/TDAs_algogram/usuario"
	"fmt"
	"strings"
)

type post struct {
	descripcion      string
	likes            TDAdic.DiccionarioOrdenado[string, string]
	usuarioPosteador usuario.Usuario
	id               int
}

func (p post) LikearPost(usuarioNombre string) string {
	if !p.likes.Pertenece(usuarioNombre) {
		p.likes.Guardar(usuarioNombre, usuarioNombre)
	}
	return "Post likeado"
}

func (p post) ObtenerPoster() usuario.Usuario {
	return p.usuarioPosteador
}

func (p post) MostrarLikes() string {
	var nombre string
	p.likes.Iterar(func(clave string, dato string) bool {
		nombre += fmt.Sprintf("\n%s", dato)
		return true
	})
	return fmt.Sprintf("El post tiene %d likes:%s", p.likes.Cantidad())
}

func (p post) MostrarPost() string {
	return fmt.Sprintf("Post ID %d\n%s\nLikes: %d", p.id, p.descripcion, p.likes.Cantidad())
}

func CrearPost(id int, descripcion string, usuarioPosteador usuario.Usuario) Post {
	p := new(post)
	p.descripcion = descripcion
	p.likes = TDAdic.CrearABB[string, string](strings.Compare)
	p.usuarioPosteador = usuarioPosteador
	p.id = id
	return p
}
