package main

import (
	"fmt"
	"runtime"

	"github.com/salernolucas93/godesdecero/arreglos_slices"
	d "github.com/salernolucas93/godesdecero/defer_panic"
	"github.com/salernolucas93/godesdecero/ejercicios"
	"github.com/salernolucas93/godesdecero/files"
	"github.com/salernolucas93/godesdecero/funciones"
	e "github.com/salernolucas93/godesdecero/interfaces_impl"
	"github.com/salernolucas93/godesdecero/iteraciones"
	"github.com/salernolucas93/godesdecero/mapas"
	m "github.com/salernolucas93/godesdecero/modelos"
	"github.com/salernolucas93/godesdecero/users"
	"github.com/salernolucas93/godesdecero/variables"
	"github.com/salernolucas93/godesdecero/webserver"
	//"github.com/salernolucas93/godesdecero/goroutines"
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

	/* Mapas */
	mapas.MostrarMapas()

	/* Estructuras */
	users.AltaUsuario()

	/* Interfaces */
	Pedro := new(m.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(m.Mujer)
	e.HumanosRespirando(Maria)

	/* Defer - Panic */
	d.VerDefer()
	//d.VerPanic()

	/* GO Routines y Channels */
	//canal1 := make(chan bool)
	//go goroutines.MiNombreLento("Lucas Salerno", canal1) // La instruccion "go" indica que esta instruccion se ejecutara de manera asincrona
	//defer func() {
	//<-canal1 // Awake - La ejecucion esperara a la asignacion del channel
	//}()
	//fmt.Println("Estoy aqui")

	/* Servidor Web */
	webserver.MiWebServer()
}
