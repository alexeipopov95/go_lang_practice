package main

import "fmt"

func __________main() {
	// diccionarios en go (maps)

	// basicamente lo que hago aca es:
	// declaro un mapa y con el metodo 'make' que sirve para crear diccionarios. Ademas de otras cosas.
	// para crear un dict ponemos map['tipo de dato que hara referencia a la key']tipo de dato que hace referencia al value
	mapa := make(map[string]int)
	mapa["Jose"] = 14
	mapa["Pedra"] = 574
	fmt.Println(mapa)

	for key, value := range mapa {
		fmt.Println(key, value)
	}
	/*
		Aca hay un WARNING. Al hacerse de forma concurrente, los valores van a venir desordenados. Eso significa que primero puede vernir 1 cosa y luego otra.
		Para arreglar eso hay que usar slices.
	*/

	// obtener uno de los valores
	value := mapa["Jose"]
	fmt.Println(value)

	// Si hay algun typo o error al momento de buscar por key
	// Go directamente te retorna el valor 0. Te pone el mismo valor que tiene el zero value.
	// para saber si esta todo bien, hay que hacer lo siguiente.
	value, ok := mapa["Jose"]
	fmt.Println(value, ok)
	// basicamente lo que te indica es un true o false, entonces vos en base a eso podes laburar con mayor seguridad.
	// los maps son mas eficientes que los slices o arrays dado que de forma nativa implementan concurrencia.
}
