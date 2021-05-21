package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Juan"] = 15

	fmt.Println(m)

	// Recorrer Map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar 1 valor. En Go si accedes a una llave que no existe te regresa 0, al escribir ok nos regresa un bool si existe o no.
	value, ok := m["Josae"]
	fmt.Println(ok, value)
}
