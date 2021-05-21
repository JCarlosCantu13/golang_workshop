package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func updateBrand(myPc *pc) {
	myPc.brand = "Acer"
}

func main() {
	fmt.Println()
	fmt.Println("Ejemplo 5 Punteros en funciones")

	myPc := pc{ram: 16, disk: 200, brand: "apple"}
	fmt.Println("Marca Actual", myPc.brand)
	updateBrand(&myPc)
	fmt.Println("Marca Actualizada", myPc.brand)
}
