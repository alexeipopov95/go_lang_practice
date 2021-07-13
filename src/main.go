package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func main() {
	// punteros y structs

	a := 50
	// el & seguido de una variable hara referencia a la referencia de valor en memoria.
	b := &a

	// si imprimimos dicha referencia entonces tenemos como resultado 0xc000012090
	fmt.Println(b)
	// Para mostrar el valor de una direccion en memoria le ponemos el * antes
	fmt.Println(*b)

	// ¿para que carajo me sirve hacer esto?
	// es como un enlace simbolico a una variable...

	*b = 100
	fmt.Println(a)
	// Aqui a me va a dar 100. ¿porque si a era 50 y b apunta a su direccion en memoria...?
	// porque comparten una misma direccion en memoria.

	// nuevamente ¿porque carajo quiero esto?
	// Se utilizan principalmente para crear funciones mas custom y para poder llevar las funciones de una libreria
	// a otro paquete de forma mas facil y eficiente.

	// ejemplo *ver arriba la construccion del struct 'pc'*

	myPc := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPc)

	myPc.ping()
}
