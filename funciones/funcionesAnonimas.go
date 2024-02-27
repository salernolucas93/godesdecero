package funciones

import "fmt"

func Calculos() {
	var numero3 int = 32
	var numero4 int = 243

	/* Funcion Anonima asociada a una variable */
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	//La funcion se puede sobreescribir en su logica (pero no en sus parametros, no es una Sobrecarga propiamente dicha)
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(10, 25))

	// Las funciones anonimas pueden tambien ser retornadas o ingresadas como parametros mediante variables
}
