package main

import "fmt"

func calcMath(x ...int) int {
	total := 0
	for _, xx := range x {
		total += xx
	}
	return total
}

func escrever(t string, n ...int) {
	for _, i := range n {
		fmt.Println(t, i)
	}
}

func main() {

	a := calcMath(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(a)

	escrever("Olá mundo", 1, 2, 3, 4, 5)

}
