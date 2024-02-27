package mapas

import "fmt"

func MostrarMapas() {
	// Declaracion de mapa con make
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	// Declaracion de Mapa sin make - Obliga a declarar el/los elementos
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato { //Recorrer un mapa
		fmt.Printf("Equipo %s tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid") // Elimina elemento del mapa
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"] // Busca si valor segun clave existe
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
