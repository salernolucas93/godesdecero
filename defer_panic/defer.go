package defer_panic

import (
	"fmt"
	"log"
)

func VerDefer() {
	/* DEFER - Define cual va a ser la ultima instruccion en ejecutarse cuando salga de la funcion */
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

func VerPanic() {
	/* PANIC - Aborta el programa cuando existe una falla, y envia un mensaje a consola */
	defer func() {
		/* RECOVER - Se usa con defer obligatoriamente. Obliga al programa a ejecutar el defer, aun si hubo un panic
		* Si no hubo panic, grabara un nil en la variable
		 */
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
