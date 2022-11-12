package procesar_archivos

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
	for s.Scan() {
		usuarioNombre := s.Text()
		nuevoUsuario := usuario.CrearUsuario(usuarioNombre)
		usuariosHash.Guardar(usuarioNombre, nuevoUsuario)
	}
	return nil
}
