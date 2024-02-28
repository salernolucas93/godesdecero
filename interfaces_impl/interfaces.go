package interfaces_impl

import (
	"fmt"

	"github.com/salernolucas93/godesdecero/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", humano.Sexo())
}
