package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MostrarSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:]   // Slice desde vector, desde posicion 3 hasta final
	porcion2 := arreglo[:5]  // Slice desde vector, desde inicio hasta posicion 5
	porcion3 := arreglo[6:7] // Slice desde vector, desde posicion 6 hasta posicion 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) // Crea Slice de 5 elementos y 20 de capacidad
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
