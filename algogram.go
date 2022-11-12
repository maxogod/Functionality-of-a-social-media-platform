package main

import (
	"algogram/errores"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ARCHIVO_USUARIOS = os.Args[1]

func main() {
	// TODO procesar archivos
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
