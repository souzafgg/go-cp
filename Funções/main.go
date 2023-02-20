package main

import (
	"fmt"
)

// func calcM(n1, n2 int8) (int8 int8) {
// 	soma := n1 + n2
// 	subr := n1 - n2
// 	return soma, subr
// }

func main() {
	x := somar(10, 15)
	fmt.Println(x)

	var f = func(txt string) {
		fmt.Println("txt")
	}

	f("Função f")

	// rsoma, rsub := calcM(20, 25)
	// fmt.Println(rsoma, rsub)

}

func somar(n, m int) int {
	return n + m
}
