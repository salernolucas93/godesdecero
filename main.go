package main

import (
	"fmt"

	"github.com/salernolucas93/godesdecero/variables"
)

func main() {
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvertirTexto(200)
	fmt.Println(estado)
	fmt.Println(texto)
}
