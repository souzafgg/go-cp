package main

import "fmt"

func main() {
	n := 10

	if n > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Menor que 10")
	}

	// if init só funciona no escopo do if
	if n2 := n; n2 > 0 {
		fmt.Println("É maior")
	} else {
		fmt.Println("É menor")
	}
}
