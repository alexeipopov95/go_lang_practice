package main

import (
	"fmt"
	"go_lang_practice/src/newpackage"
)

func ___________main() {
	var myPublicCar newpackage.CarPublic
	myPublicCar.Brand = "Ferrari"
	myPublicCar.Year = 2020
	fmt.Println((myPublicCar))

	// No se puede. Es privado.
	//var myPrivateCar newpackage.carPrivate
	//fmt.Println(myPrivateCar)

	// aca me traje una funcion publica que no hace nada
	newpackage.PrintMessage()

	// aca me traje una funcion publica que incluye componentes privados.
	newpackage.Lawea()

}
