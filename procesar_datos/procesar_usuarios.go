package procesar_datos

import (
	diccionario "algogram/TDAs/hash"
	"algogram/TDAs_algogram/usuario"
	"algogram/errores"
	"bufio"
	"os"
)

func ObtenerTodosUsuarios(args []string, usuariosHash diccionario.Diccionario[string, usuario.Usuario]) error {
	if len(args) != 1 {
		return errores.ErrorLecturaArchivos{}
	}
	archivo, err := os.Open(args[0])
	if err != nil {
		return errores.ErrorLecturaArchivos{}
	}
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	for i := 0; s.Scan(); i++ {
		usuarioNombre := s.Text()
		nuevoUsuario := usuario.CrearUsuario(usuarioNombre, i)
		usuariosHash.Guardar(usuarioNombre, nuevoUsuario)
	}
	return nil
}
