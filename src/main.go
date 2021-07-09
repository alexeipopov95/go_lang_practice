// Para el momento de desarrollo y si queres testearlo, tendras que correr 'go run main.go'
// para compilarlo hay que hacer primero un build 'go build main.go' y recien despues ejecutarlo con ./main
// es mas eficiente compilar y luego correr.

package main

import "fmt"

func main() {
	// fmt es el paquete que permite administrar los inputs y outputs de la termianl.
	// Hace falta importarlo en caso de no existir.
	// fmt.Println -> es el print() de python.

	//fmt.Println("Hola wey")

	// constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(pi)
	fmt.Println(pi2)

	// variables enteras
	// el : indica que una variable no ha sido creada anteriormente.
	// entonces este la crea y le asigna el valor.
	// si le paso sin los : significa que antes tuve que crear la variable.
	base := 12

	// otra forma de crear una variable
	var altura int = 14

	// podemos declarar una variable sin instanciarla
	var area int

	fmt.Println(base, altura, area)

	// zero values. Son valores por defecto, como en python (None).
	// WARNING: NO son valores nulos, sino por defecto.
	var zvalue1 int
	var zvalue2 float64
	var zvalue3 string
	var zvalue4 bool

	fmt.Println("Zero value 1:", zvalue1) // nulo para int es un 0
	fmt.Println("Zero value 2:", zvalue2) // nulo para float es un 0
	fmt.Println("Zero value 3:", zvalue3) // nulo para string es un string "vacio" NO nulo.
	fmt.Println("Zero value 4:", zvalue4) // nulo para bool es false.

	// calcular el area de un cuadrado.
	const baseCuadrado int = 10

	// TODO: El := actúa como comodín al momento de asignarle un valor a una variable?
	//		 Osea que no le tengo que especificar que es un tipo de var, const nisiquiera el tipo?
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area de un cuadrado:", areaCuadrado)

	// Operadores aritmeticos

	x := 10
	y := 50

	// SUMA
	result := x + y
	fmt.Println("Resultado de la suma:", result)

	// RESTA (negativos)
	result = x - y
	fmt.Println("Resultado de la resta:", result)

	// RESTA (positivos)
	result = y - x
	fmt.Println("Resultado de la resta:", result)

	// MULTIPLICACION
	result = x * y
	fmt.Println("Resultado de la multiplicacion:", result)

	// DIVISION
	result = x / y
	fmt.Println("Resultado de la division:", result)

	// DIVISION
	result = y / x
	fmt.Println("Resultado de la division:", result)

	// Modulo
	result = y % x
	fmt.Println("Resultado de la division modulo entre y y x:", result)

	// Modulo
	result = x % y
	fmt.Println("Resultado de la division modulo entre x y y:", result)

	// incremental
	x++
	y++
	fmt.Println("Incremental de X:", x)
	fmt.Println("Incremental de Y:", y)

	// Decremental
	x--
	y--
	fmt.Println("Decremental de X:", x)
	fmt.Println("Decremental de Y:", y)

	// Tipos de datos primitivos o tipo de dato
	// Esto es una forma de declarar enteros positivos y negativos para diferentes "arquitecturas"
	// int/int8/int16/int32/int64 -> int solo, toma el tamaño de bits en el que esta basado el procesador.

	// Esto es una forma de declarar enteros POSITIVOS para diferentes arquitecturas
	// uint/uint8/uint/16/uint32/uint64 -> uint solo, toma el tamaño de bits en el que esta basado el procesador.

	// float32 para 32 bits con ambos signos.
	// float64 para 64 bits con ambos signos.

	// string se deben de declarar con comillas dobles.
	// bool es true o false

	// Declarando unas variables para laburar mejor el 'fmt'
	helloMessage := "Hello"
	worldMessage := "World"

	// Println -> agrega salto de linea de forma automatica.
	fmt.Println(helloMessage, worldMessage)

	// printf -> Permite agregar funcionalidad extra. Es el .format() o el f"" de python.
	// Pero no discrimina el tipo de valor.
	// %s es string
	// %d es un entero?
	// %v es variable. Basicamente es mas generico. Pero mejor especificar que tipo de variable va a ser como buena practica.
	name := "Alex"
	number := 500
	fmt.Printf("Name %s and number is %d\n", name, number)
	// Mala practica

	fmt.Printf("Name %v and number is %v\n", name, number)

	// Sprintf -> este guarda en una variable el printf para luego ser lanzado en otra ocasion.
	message := fmt.Sprintf("Name %s and number is %d\n", name, number)
	fmt.Println(message)

	// Paquete para conocer el tipo de dato
	fmt.Printf("name: %T\n", name)
	fmt.Printf("number: %T\n", number)

}
