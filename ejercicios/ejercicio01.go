package ejercicios

import (
	"strconv"
)

func Conversor(parametro string) (int, string) {
	numero, _ := strconv.Atoi(parametro)
	var texto string
	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}
	return numero, texto
}
