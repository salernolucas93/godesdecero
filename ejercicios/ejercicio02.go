package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func MultiplicarTabla() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Ingrese numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}

	return texto
}
