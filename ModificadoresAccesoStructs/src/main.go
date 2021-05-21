package main

import (
	"ModificadoresAccesoStructs/src/mypackage"
	"fmt"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Kia"
	myCar.Year = 2021

	fmt.Println(myCar)
}
