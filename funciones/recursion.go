package funciones

import "fmt"

func Exponencia(valor int) {
	/* Recursion - Funcion que se llama a si misma */
	if valor > 100000000 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor * 2)
}
