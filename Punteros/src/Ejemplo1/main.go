package main

import "fmt"

func main() {
	fmt.Println("test")

	fmt.Println("Como declarar un puntero")

	v := 10

	var p *int
	fmt.Printf("%T", p) // %T representation of the type of the value
	fmt.Println()

	var p1 = new(int)
	fmt.Printf("%T", p1)
	fmt.Println()

	p2 := &v
	fmt.Printf("%T", p2)
	fmt.Println()
}
