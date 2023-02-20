package main

import "fmt"

func main() {

	a := 10
	b := a

	fmt.Println(a, b)

	a++

	fmt.Println(a, b)

	// Ponteiro

	var c int
	var p *int

	c = 100
	p = &c

	fmt.Println(c, *p) //Desreferenciação

}
