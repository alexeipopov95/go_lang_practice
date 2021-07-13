package newpackage

import "fmt"

// Basicamente si yo quiero que mi struct sea utilizada por el resto de las cosas
// lo que debo hacer es que la struct comience con mayuscula y que cada atributo que quiero
// mostrar este tambien comenzando con mayuscula.

// 'nombre_clase' 'descripcion'
// CarPublic Car con acceso publico.
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMEssage Imprime un mensaje aleatorio
func PrintMessage() {
	fmt.Println("Imprimo mensaje.")
}

func Lawea() {
	var myCar carPrivate
	myCar.brand = "BLERG"
	fmt.Println(myCar)
}
