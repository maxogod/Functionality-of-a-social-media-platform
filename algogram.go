package main

import (
	diccionario "algogram/TDAs/hash"
	"algogram/TDAs_algogram/post"
	"algogram/TDAs_algogram/usuario"
	"algogram/errores"
	"algogram/procesar_archivos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ARGS = os.Args[1:]

func main() {
	var (
		usuarios                 = diccionario.CrearHash[string, usuario.Usuario]()
		posts                    = diccionario.CrearHash[int, post.Post]()
		logueado usuario.Usuario = nil
	)
	errDeProcesamiento := procesar_archivos.ObtenerTodosUsuarios(ARGS, usuarios)
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
			nombreUsuario := entrada[1]
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

		case "ver_siguiente_feed":
			if logueado != nil {
				siguientePost, err := logueado.VerSigPost()
				if err == nil {
					fmt.Println(siguientePost)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(errores.ErrorVerPost{})
			}

		case "likear_post":
		case "mostrar_likes":
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
}
