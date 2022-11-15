package main

import (
	diccionario "algogram/TDAs/hash"
	"algogram/TDAs_algogram/post"
	"algogram/TDAs_algogram/usuario"
	"algogram/errores"
	"algogram/procesar_datos"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ARGS = os.Args[1:]

func main() {
	var (
		usuarios                 = diccionario.CrearHash[string, usuario.Usuario]()
		posts                    = diccionario.CrearHash[int, post.Post]()
		logueado usuario.Usuario = nil
	)
	errDeProcesamiento := procesar_datos.ObtenerTodosUsuarios(ARGS, usuarios)
	if errDeProcesamiento != nil {
		fmt.Println(errDeProcesamiento)
		return
	}
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		entrada := strings.Split(scan.Text(), " ")
		comando := entrada[0]

		switch comando {
		case "login":
			nombreUsuario := strings.Join(entrada[1:], " ")
			if usuarios.Pertenece(nombreUsuario) && logueado == nil {
				logueado = usuarios.Obtener(nombreUsuario)
				fmt.Println(fmt.Sprintf("Hola %s", logueado.ObtenerNombre()))
			} else if logueado != nil {
				fmt.Println(errores.ErrorLogin{})
			} else {
				fmt.Println(errores.ErrorUsuarioInvalido{})
			}

		case "logout":
			if logueado != nil {
				logueado = nil
				fmt.Println("Adios")
			} else {
				fmt.Println(errores.ErrorSinUsuarioLogueado{})
			}

		case "publicar":
			if logueado != nil {
				mensaje := strings.Join(entrada[1:], " ")
				nuevoPost := post.CrearPost(posts.Cantidad(), mensaje, logueado.ObtenerId(), logueado.ObtenerNombre())
				posts.Guardar(posts.Cantidad(), nuevoPost)
				procesar_datos.GuardarPostEnFeeds(nuevoPost, usuarios)
				fmt.Println("Post publicado")
			} else {
				fmt.Println(errores.ErrorSinUsuarioLogueado{})
			}

		case "ver_siguiente_feed":
			if logueado != nil {
				idPost, err := logueado.VerSigPost()
				if err == nil {
					fmt.Println(posts.Obtener(idPost).MostrarPost())
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(errores.ErrorVerPost{})
			}

		case "likear_post":
			postId, _ := strconv.Atoi(entrada[1])
			if logueado != nil && posts.Pertenece(postId) {
				postActual := posts.Obtener(postId)
				postActual.LikearPost(logueado.ObtenerNombre())
				fmt.Println("Post likeado")
			} else {
				fmt.Println(errores.ErrorDarLike{})
			}

		case "mostrar_likes":
			postId, _ := strconv.Atoi(entrada[1])
			if posts.Pertenece(postId) {
				mensajeLikes, err := posts.Obtener(postId).MostrarLikes()
				if err == nil {
					fmt.Println(mensajeLikes)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(errores.ErrorVerLike{})
			}

		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
}
