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
	fmt.Println("Ejemplo 4 valores por referencia en metodos")

	myPc := pc{ram: 16, disk: 200, brand: "apple"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println("Estado actual", myPc.ram)
	myPc.duplicateRAM()
	fmt.Println("Estado despues de duplicar", myPc.ram)

}
