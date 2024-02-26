package ejercicios

import (
	"strconv"
)

func Conversor(parametro string) (int, string) {
	numero, err := strconv.Atoi(parametro)
	if err != nil { // Agregado Manejo de Error (En Go lo idea es no omitirlo)
		return 0, "Hubo un error: " + err.Error()
	}
	var texto string
	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}
	return numero, texto
}
