package main

import (
	diccionario "algogram/TDAs/hash"
	"algogram/TDAs_algogram/post"
	"algogram/TDAs_algogram/procesar_archivos"
	"algogram/TDAs_algogram/usuario"
	"algogram/errores"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ARGS = os.Args[1:]

func main() {
	var (
		usuarios = diccionario.CrearHash[string, usuario.Usuario]()
		posts    = diccionario.CrearHash[int, post.Post]()
	)
	err := procesar_archivos.ObtenerTodosUsuarios(ARGS, usuarios)
	if err != nil {
		print(err.Error())
		return
	}
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		entrada := strings.Split(scan.Text(), " ")
		comando := entrada[0]

		switch comando {
		case "login":
		case "logout":
		case "publicar":
		case "ver_siguiente_feed":
		case "likear_post":
		case "mostrar_likes":
		default:
			fmt.Println(new(errores.ErrorParametros))
		}
	}
}
