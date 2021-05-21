package main

import (
	"fmt"

	"github.com/JCarlosCantu13/golang_workshop/ModificadoresAccessoStructs/src/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Kia"
	myCar.Year = 2021

	fmt.Println(myCar)
}
