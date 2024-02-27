package main

import (
	"fmt"
	"runtime"

	"github.com/salernolucas93/godesdecero/arreglos_slices"
	"github.com/salernolucas93/godesdecero/ejercicios"
	"github.com/salernolucas93/godesdecero/files"
	"github.com/salernolucas93/godesdecero/funciones"
	"github.com/salernolucas93/godesdecero/iteraciones"
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
	numero, mensaje := ejercicios.Conversor("2000") // Mayor
	fmt.Println(numero)
	fmt.Println(mensaje)

	numero, mensaje = ejercicios.Conversor("50") //Menor
	fmt.Println(numero)
	fmt.Println(mensaje)

	numero, mensaje = ejercicios.Conversor("aaaaa") // Manejo de Error
	fmt.Println(numero)
	fmt.Println(mensaje)

	/* Mostrar y Aceptar Datos */
	//teclado.IngresoNumeros()

	/* Ciclos */
	iteraciones.Iterar()

	/* Ejercicio 02 (Modificado para MAnejo de Archivos) */
	//fmt.Println(ejercicios.MultiplicarTabla())

	//files.GrabarTabla()
	//files.SumarTabla()

	files.LeerArchivo()

	/* Funciones Anonimas y Closures */
	//funciones.Calculos()
	//funciones.LlamarClosure()

	/* Recursion */
	funciones.Exponencia(2)

	/* Arreglos y Slices */
	arreglos_slices.MostrarArreglos()
	arreglos_slices.MostrarSlice()
	arreglos_slices.Capacidad()
}
