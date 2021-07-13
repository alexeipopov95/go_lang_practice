package main

import "fmt"

func _______main() {
	// Array en GO son INMUTABLES.
	var array [4]int

	// Editar por elementos
	array[0] = 1
	array[1] = 2

	fmt.Println(array)

	// funciones importantes en GO
	// len() dice cuantos elementos hay en un array
	fmt.Println("Longitud del array:", len(array))

	// cap() dice la capacidad maxima del array
	fmt.Println("Capacidad maxima del array:", cap(array))

	// SLICE
	// Es similar a los array, pero no se indican cuantos elementos pueden haber.
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println("Longitud del slice:", len(slice))
	fmt.Println("Capacidad maxima del slice:", cap(slice))

	// Slicing
	// Se usa en los arrays para interactuar con los elementos.
	//Metodos en el slice IGUAL QUE EN PYTHON
	fmt.Println(slice[0])   // imprime el elemento 0
	fmt.Println(slice[:3])  // imprime todo HASTA el elemento 3
	fmt.Println(slice[2:4]) // imprime unicamente aquellos elementos entre el 2 y el 4
	fmt.Println(slice[4:])  // imprime todo desde el elemento 4 en adelante

	// Append
	// agregar elemento en un slice
	slice = append(slice, 7)
	fmt.Println(slice)

	// Agregar una lista
	newSlide := []int{8, 9, 10}
	slice = append(slice, newSlide...) // estos ... puntos basicamente lo que hace es itera internamente la lista
	// y desempaqueta la lista nueva en la actual.
	fmt.Println(slice)

}
