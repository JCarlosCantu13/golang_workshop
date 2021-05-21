package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram *= 2
}

func main() {

	fmt.Println("Como declarar un puntero")
	v := 10

	var p *int
	fmt.Printf("%T", p) // %T representation of the type of the value

	var p1 = new(int)
	fmt.Printf("%T", p1)

	p2 := &v
	fmt.Printf("%T", p2)

	fmt.Println()
	fmt.Println("Ejemplo 2 Referencia & ")
	var valor int = 5
	fmt.Println("La direccion de memoria de valor es: ", &valor)
	fmt.Println("El valor es: ", valor)
	fmt.Printf("El valor es: %d como vez?", valor)

	fmt.Println()
	fmt.Println("Ejemplo 3 Desreferenciar ")
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

	fmt.Println()
	fmt.Println("Ejemplo 4")
	myPc := pc{ram: 16, disk: 200, brand: "apple"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println("Estado actual", myPc.ram)
	myPc.duplicateRAM()
	fmt.Println("Estado despues de duplicar", myPc.ram)
}
