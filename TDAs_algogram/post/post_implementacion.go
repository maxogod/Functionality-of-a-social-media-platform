package post

import (
	TDAdic "algogram/TDAs/abb"
	"fmt"
	"strings"
)

type PostImplementation struct {
	descripcion        string
	likes              TDAdic.DiccionarioOrdenado[string, string]
	idUsuarioPosteador int
	id                 int
}

func (p PostImplementation) LikearPost(usuarioNombre string) string {
	if !p.likes.Pertenece(usuarioNombre) {
		p.likes.Guardar(usuarioNombre, usuarioNombre)
	}
	return "Post likeado"
}

func (p PostImplementation) ObtenerPosterID() int {
	return p.idUsuarioPosteador
}

func (p PostImplementation) ObtenerPostID() int {
	return p.id
}

func (p PostImplementation) MostrarLikes() string {
	var nombre string
	p.likes.Iterar(func(clave string, dato string) bool {
		nombre += fmt.Sprintf("\n%s", dato)
		return true
	})
	return fmt.Sprintf("El PostImplementation tiene %d likes:%s", p.likes.Cantidad())
}

func (p PostImplementation) MostrarPost() string {
	return fmt.Sprintf("Post ID %d\n%s\nLikes: %d", p.id, p.descripcion, p.likes.Cantidad())
}

func CrearPost(id int, descripcion string, idUsuarioPosteador int) Post {
	p := new(PostImplementation)
	p.descripcion = descripcion
	p.likes = TDAdic.CrearABB[string, string](strings.Compare)
	p.idUsuarioPosteador = idUsuarioPosteador
	p.id = id
	return p
}
