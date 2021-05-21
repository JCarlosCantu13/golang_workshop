package main

import "fmt"

func main() {
	fmt.Println("Ejemplo 2 Referencia & ")

	var valor int = 5
	fmt.Println("La direccion de memoria de valor es: ", &valor)
	fmt.Println("El valor es: ", valor)
	fmt.Printf("El valor es: %d como vez?", valor)

	fmt.Println()
	fmt.Println("Ejemplo 3 Desreferenciar *")

	a := 50
	b := &a // A la variable b le asignamos la dirección de memoria(referencia) de a

	fmt.Printf("La variable a es: %d \n", a)
	fmt.Printf("La dirección de memoria de a es: %v \n", &a)
	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n", b)
	fmt.Printf("Al desrefenciar el puntero p obtengo el valor: %d \n", *b)

	fmt.Println(a, b)
	fmt.Println(*b) // tomar en cuenta que & es para acceder a la dirección de memoria y *b es para acceder al valor almacenado en esa memoria

	*b = 100 // en este caso modificamos el valor de b, pero b va a afectar la dirección de memoria y cambiará el valor de a
	fmt.Println(a)
}
