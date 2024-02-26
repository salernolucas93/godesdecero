package main

import (
	"fmt"
	"runtime"

	"github.com/salernolucas93/godesdecero/ejercicios"
	"github.com/salernolucas93/godesdecero/variables"
)

func main() {
	/* Variables */
	variables.MostrarEnteros()

	/* Resto de Variables */
	variables.RestoVariables()

	/* Funciones */
	estado, texto := variables.ConvertirTexto(200)
	fmt.Println(estado)
	fmt.Println(texto)

	/* Condiciones */
	if os := runtime.GOOS; os == "OS X." {
		fmt.Println("Esto no es Linux, es ", os)
	} else {
		fmt.Println("Esto es Linux")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux: ")
	case "darwin":
		fmt.Println("Esto es Darwin: ")
	default:
		fmt.Printf("%s \n", os)
	}

	/* Ejercicio 01 */
	numero, mensaje := ejercicios.Conversor("2000")
	fmt.Println(numero)
	fmt.Println(mensaje)

	numero, mensaje = ejercicios.Conversor("50")
	fmt.Println(numero)
	fmt.Println(mensaje)
}
