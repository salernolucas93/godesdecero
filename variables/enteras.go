package variables

import "fmt"

func MostrarEnteros() {
	var intcomun int      // Variable por declaracion
	intde32 := int32(10)  // Variable por asignacion (dos puntos e igual := para asignar)
	intde64 := int64(100) // La funcion int64 convierte el valor a int64 (idem int32).

	fmt.Println("intcomun: ", intcomun)
	fmt.Println("intde32: ", intde32)
	fmt.Println("intde64: ", intde64)
}
