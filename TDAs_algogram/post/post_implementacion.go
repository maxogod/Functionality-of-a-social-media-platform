package post

import (
	TDAdic "algogram/TDAs/abb"
	"algogram/errores"
	"fmt"
	"strings"
)

type postImplementation struct {
	descripcion        string
	likes              TDAdic.DiccionarioOrdenado[string, string]
	idUsuarioPosteador int
	id                 int
	nombrePoster       string
}

func (p postImplementation) LikearPost(usuarioNombre string) string {
	if !p.likes.Pertenece(usuarioNombre) {
		p.likes.Guardar(usuarioNombre, usuarioNombre)
	}
	return "Post likeado"
}

func (p postImplementation) ObtenerPosterID() int {
	return p.idUsuarioPosteador
}

func (p postImplementation) ObtenerPostID() int {
	return p.id
}

func (p postImplementation) MostrarLikes() (string, error) {
	if p.likes.Cantidad() == 0 {
		return "", errores.ErrorVerLike{}
	}
	var nombre string
	p.likes.Iterar(func(clave string, dato string) bool {
		nombre += fmt.Sprintf("\n\t%s", dato)
		return true
	})
	return fmt.Sprintf("El post tiene %d likes:%s", p.likes.Cantidad(), nombre), nil
}

func (p postImplementation) MostrarPost() string {
	return fmt.Sprintf("Post ID %d\n%s dijo: %s\nLikes: %d", p.id, p.nombrePoster, p.descripcion, p.likes.Cantidad())
}

func CrearPost(id int, descripcion string, idUsuarioPosteador int, nombrePosteador string) Post {
	p := new(postImplementation)
	p.descripcion = descripcion
	p.likes = TDAdic.CrearABB[string, string](strings.Compare)
	p.idUsuarioPosteador = idUsuarioPosteador
	p.id = id
	p.nombrePoster = nombrePosteador
	return p
}
