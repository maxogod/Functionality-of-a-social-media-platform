package diccionario_test

import (
	TDAdic "diccionario/abb"
	"diccionario/errores"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
)

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDAdic.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.EqualValues(t, 10, dic.Borrar("A"))
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDAdic.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDAdic.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDAdic.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	err := new(errores.ErrorNoEncontrado)

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, err.Error(), func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, err.Error(), func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, err.Error(), func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, err.Error(), func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, err.Error(), func() { dic.Obtener(claves[1]) })
}

func TestBorrarUnaHoja(t *testing.T) {
	t.Log("Test borrar hoja de un arbol")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	clave1 := "A"
	clave2 := "B" // A borrar
	clave3 := "C"
	valor1 := 1
	valor2 := 0 // A borrar
	valor3 := 2
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Borrar(clave3)
	require.False(t, dic.Pertenece(clave3))

	// Checkea que el arbol quedo como se supone
	i := valor1
	dic.Iterar(func(clave string, valor int) bool {
		require.EqualValues(t, i, valor)
		i++
		return true
	})

}

func TestBorrarNodoConHijo(t *testing.T) {
	t.Log("Test borrar nodo con un hijo")
	dic := TDAdic.CrearABB[string, int](strings.Compare)

	clave1 := "A"
	clave2 := "B" // A borrar
	clave3 := "C"
	valor1 := 0
	valor2 := 1 // A borrar
	valor3 := 2
	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	dic.Borrar(clave3)
	require.False(t, dic.Pertenece(clave3))

	// Checkea que el arbol quedo como se supone
	i := valor1
	dic.Iterar(func(clave string, valor int) bool {
		require.EqualValues(t, i, valor)
		i++
		return true
	})
}

func TestBorrarNodoDosHijos(t *testing.T) {
	t.Log("Creamos un arbol con claves numericas donde la raiz tiene 2 hijos, y la borramos, despues se checkea que " +
		"el arbol quedo como esperado.")
	dic := TDAdic.CrearABB[int, string](func(x, y int) int {
		if x < y {
			return -1
		} else if x > y {
			return 1
		}
		return 0
	})
	clave1 := 2
	clave2 := 1
	clave3 := 3
	valor1 := "Gatito"
	valor2 := "Perrito"
	valor3 := "Pecesito"

	dic.Guardar(clave1, valor1)
	dic.Guardar(clave2, valor2)
	dic.Guardar(clave3, valor3)
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(clave1) && dic.Pertenece(clave2) && dic.Pertenece(clave3))
	require.EqualValues(t, valor1, dic.Obtener(clave1))
	require.EqualValues(t, valor2, dic.Obtener(clave2))
	require.EqualValues(t, valor3, dic.Obtener(clave3))
	require.EqualValues(t, valor1, dic.Borrar(clave1))
	require.False(t, dic.Pertenece(clave1))

	// Checkeamos que despues de borrar la raiz el arbol quedo como se supone.
	i := clave2 // = 1
	dic.Iterar(func(clave int, valor string) bool {
		require.EqualValues(t, i, clave)
		i += 2 // = 3
		return true
	})
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}
	f := func(a1, a2 avanzado) int {
		return strings.Compare(a1.x.a, a2.x.a)
	}
	dic := TDAdic.CrearABB[avanzado, int](f)

	a1 := avanzado{w: 10, z: "que", x: basico{a: "onda", b: 8}, y: basico{a: "gente", b: 10}}
	a2 := avanzado{w: 10, z: "piola", x: basico{a: "ahre", b: 14}, y: basico{a: "epico", b: 5}}
	a3 := avanzado{w: 10, z: "sant1", x: basico{a: "max0", b: 8}, y: basico{a: "cracks", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDAdic.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestIter(t *testing.T) {
	t.Log("Test con iterador interno sin rango")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	dic.Guardar("D", 4)
	dic.Guardar("C", 3)
	dic.Guardar("B", 2)
	dic.Guardar("F", 6)
	dic.Guardar("E", 5)
	dic.Guardar("G", 7)
	dic.Guardar("A", 1)

	// Iterar hasta el final
	const valorUltimaClave = 7
	i := 1
	dic.Iterar(func(clave string, dato int) bool {
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, valorUltimaClave+1, i)

	// Iterar hasta Condicion
	const valorUltimaClaveVista = 5
	i = 1
	dic.Iterar(func(clave string, dato int) bool {
		if strings.Compare(clave, "F") >= 0 {
			return false
		}
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, valorUltimaClaveVista+1, i)
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterRango(t *testing.T) {
	t.Log("Test de iterador interno con rango")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	dic.Guardar("D", 4)
	dic.Guardar("C", 3)
	dic.Guardar("B", 2)
	dic.Guardar("F", 6)
	dic.Guardar("E", 5)
	dic.Guardar("G", 7)
	dic.Guardar("A", 1)

	var (
		desde = "B"
		hasta = "F"
	)

	// Tod0 el rango
	const valorDelHasta = 6
	i := 2 // Inicializado <i> en el valor de la clave <desde>
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, valorDelHasta+1, i)

	// Iterar rango hasta condicion
	const valorUltimoVisto = 4
	i = 2
	dic.IterarRango(&desde, &hasta, func(clave string, dato int) bool {
		if strings.Compare(clave, "E") >= 0 {
			return false
		}
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, valorUltimoVisto+1, i)

	// Iterar sin rango
	const valorUltimo = 7
	i = 1
	dic.IterarRango(nil, nil, func(clave string, dato int) bool {
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, valorUltimo+1, i)
}

func TestIterador(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	dic.Guardar("D", 4)
	dic.Guardar("C", 3)
	dic.Guardar("B", 2)
	dic.Guardar("F", 6)
	dic.Guardar("E", 5)
	dic.Guardar("G", 7)
	dic.Guardar("A", 1)

	iter := dic.Iterador()
	letras := []string{"A", "B", "C", "D", "E", "F", "G"}

	for i, letra := range letras {
		clave, dato := iter.VerActual()
		require.EqualValues(t, letra, clave)
		require.EqualValues(t, i+1, dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, letra, iter.Siguiente())
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, errores.ErrorIterTermino{}.Error(), func() { iter.Siguiente() })

}

func TestIteradorRango(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	dic := TDAdic.CrearABB[string, int](strings.Compare)
	dic.Guardar("D", 4)
	dic.Guardar("C", 3)
	dic.Guardar("B", 2)
	dic.Guardar("F", 6)
	dic.Guardar("E", 5)
	dic.Guardar("G", 7)
	dic.Guardar("A", 1)
	var desde = "C"
	var hasta = "F"
	iter := dic.IteradorRango(&desde, &hasta)
	letras := []string{"C", "D", "E", "F"}

	for i, letra := range letras {
		clave, dato := iter.VerActual()
		require.EqualValues(t, letra, clave)
		require.EqualValues(t, i+3, dato)
		require.True(t, iter.HaySiguiente())
		require.EqualValues(t, letra, iter.Siguiente())
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, errores.ErrorIterTermino{}.Error(), func() { iter.Siguiente() })
}

func TestVolumen(t *testing.T) {
	t.Log("Pruebas con un gran volumen de datos")
	dic := TDAdic.CrearABB[int, int](func(x, y int) int {
		if x < y {
			return -1
		} else if x > y {
			return 1
		}
		return 0
	})

	min := 0
	max := 1001
	for i := 0; i < 1000; i++ {
		random := rand.Intn(max-min) + min
		dic.Guardar(random, random)
	}
	dic.Guardar(500, 500)
	require.True(t, dic.Pertenece(500))
	require.EqualValues(t, 1001, dic.Cantidad())
	dic.Iterar(func(clave, valor int) bool {
		return true
	})
}
