package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/salernolucas93/godesdecero/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.MultiplicarTabla()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo: " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumarTabla() {
	var texto string = ejercicios.MultiplicarTabla()
	if !concatenar(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func concatenar(texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el Write String: " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() {
	//archivo, err := ioutil.ReadFile(fileName) // Deprecado
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
