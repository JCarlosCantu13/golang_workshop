package main

import "fmt"

func incrementar(num int) {
	num++
}

func main() {
	v := 10
	fmt.Println(v)
	incrementar(v)
	fmt.Println(v)
}
