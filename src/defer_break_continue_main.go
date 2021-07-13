package main

import "fmt"

func ______main() {
	// Defer Ejecuta la ultima funcion antes de que todo muera (termine todo el proceso de la func). Si abrimos una conexion con la base de datos
	// le metemos el defer de que cierre la conexion. Idem con archivos.
	// se puede usar mas de 1 defer en una funcion, pero la buena practica es solamente uno.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("I es", i)
			continue // este se usa cuando se cumple una condicion dada dentro del ciclo for y puede ser algo
			// que me interesa que continue.
		}

		// Break este es basicamente que cuando encuentro x cosa quiero que muera ahi.
		if i == 8 {
			fmt.Println("Break en:", i)
			break
		}

	}

}
