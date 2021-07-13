package main

import "fmt"

// aca se veran lo que son las clases pero de forma GO
// en go las class de python son structs o estructuras que tecnicamente son muy eficientes.

//arrancamos especificando type

type car struct {
	// atributos
	brand string
	year  int
}

func _________main() {

	// Asi se instancia una struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra forma de instanciar es
	var myCar2 car
	myCar2.brand = "Ferrari"
	fmt.Println(myCar2)

	// modificadores de acceso
	// si la letra con la que empieza es MAYUSCULA es publico, sino es PRIVADO.
	// laburando con newpackage y newpackage.go

}
