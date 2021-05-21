package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Si es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	// Indice y Valor
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// Solo Indice
	for i := range slice {
		fmt.Println(i)
	}

	// Solo Valor
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// Ejemplo Palindromo
	// amor a roma la lectura de derecha a izquierda como de izquierda a derecha es lo mismo
	fmt.Println("Teclea una palabra y te diremos si es palindromo o no")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	palabra := scanner.Text()
	minus := strings.ToLower(palabra)
	isPalindromo(minus)
}
