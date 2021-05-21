package main

import "fmt"

func main() {
	// Array
	var array [4]int
	array[0] = 1
	array[1] = 0
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice "slicing"
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append solo con slice porque un array son inmutables(no puede ser cambiado) y los slices son mutables(puede ser alterado o cambiado)
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append List
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // los 3 puntos indica que el slice lo va a descomprimir y agregarse de forma independiente
	fmt.Println(slice)
}
