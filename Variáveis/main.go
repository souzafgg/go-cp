package main

import "fmt"

func main() {

	var x string = "var"
	c := "var2"

	fmt.Println(x, c)

	var (
		a string = "texto"
		b string = "palavra"
		d int    = 2
	)

	fmt.Println(a, b, d)

	z, t, y := 1, 2, 3

	fmt.Println(z, t, y)

	const constante string = "Constante"

}
