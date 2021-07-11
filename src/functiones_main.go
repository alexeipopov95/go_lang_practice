package main

import "fmt"

// En go las funciones si o si son con camelCase
// si necesitamos que tenga mas de una entrada, como por ejemplo, con 3 parametros
func multiParamFunc(a, b int, c string) {
	fmt.Printf(
		"Param 1: %d\nParam 2: %d\nParam 3: %s", a, b, c)
}

// function with a return value
func returnValueFunction(a int) int {
	return a * 2
}

// function with a double return value
func doubleReturnFunction(a int) (a_1, a_2 int) {
	return a, a * 2
}

func __main() {

	//Multiple params function
	fmt.Println("Function multiParamFunc() result!")
	multiParamFunc(1, 2, "Wey!")

	fmt.Println("")
	fmt.Println("")
	// double return (like a tuple or list comprehenssion)
	fmt.Println("Function doubleReturnFunction() result!")
	result_1, result_2 := doubleReturnFunction(5)
	fmt.Printf(
		"Result 1 from doubleReturnFunction(): %d\nResult 2 from doubleReturnFunction(): %d\n",
		result_1, result_2,
	)

	// DATO UTIL
	// si tengo un valor que una funcion retorne mas de un valor y no lo quiero usar, debo declar la variable con un '_'

}
