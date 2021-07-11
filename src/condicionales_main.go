package main

import "fmt"

// CONDICIONALES!

func _____main() {

	const value_1 int = 5
	const value_2 int = 10
	const value_3 int = 15
	const value_4 int = 20

	if value_1 > value_2 {
		fmt.Printf("Value %d is higer than %d\n", value_1, value_2)
	} else {
		fmt.Printf("Value %d is higer than %d\n", value_2, value_1)
	}

	if value_2 > value_1 && value_1 < value_3 {
		fmt.Println("true")
	}

	fmt.Println(value_1 < value_2)

	// SWITCH
	// forma tradicional
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// forma mas "comun"
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par - alternativa")
	default:
		fmt.Println("Es impar - alternativa")
	}

	// switch sin condicion. Este se usa cuando hayan multiples condiciones
	value := 120
	switch {
	case value > 100 && value < 200:
		fmt.Println("Es mayor a 100 y menor que 200")
	case value < 200 && value > 50:
		fmt.Println("Es menor a 200 y mayor que 50")
	default:
		fmt.Println("No coincide")
	}

}
