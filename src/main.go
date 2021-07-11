package main

import "fmt"

func main() {
	// Defer Ejecuta la ultima funcion antes de que todo muera. Si abrimos una conexion con la base de datos
	// le metemos el defer de que cierre la conexion. Idem con archivos.
	// se puede usar mas de 1 defer en una funcion, pero la buena practica es solamente uno.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue Break
}
