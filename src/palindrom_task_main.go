package main

import (
	"fmt"
	"strings"
)

// veremos que onda si es palindromo

func reverseText(text string) string {
	var reversed string

	if len(text) == 0 {
		return text
	}

	for i := len(text) - 1; i >= 0; i-- {
		reversed += string(text[i])
	}
	return strings.ToLower(reversed)
}

func isPalindom(text string) bool {
	var isPalindrom bool = false
	reversed := reverseText(text)
	if text == reversed {
		isPalindrom = true
	}
	return isPalindrom
}

func ________main() {
	// Como recorrer un array o slice

	slice := []string{"Hola", "Como", "estas", "?"}

	// forma de iterar... en python esto seria una analogia con enumerate() donde tenes el i=indice y valor=elemento del array
	// se puede omitir tambien el indice convirtiendolo como un _
	for i, valor := range slice {
		fmt.Println("Indice:", i, "\nValor:", valor)
	}

	// determinar si una palabra es palindromo.
	var myCustomWord string = "ama"
	fmt.Println("La palabara:", myCustomWord, "\n¿Es palindromo?", isPalindom(myCustomWord))

}
