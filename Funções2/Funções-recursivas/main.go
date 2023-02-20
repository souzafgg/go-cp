package main

import "fmt"

func main() {

	// fibonacci 1 1 2 3 5 8 13

	x := uint(5)

	for i := uint(0); i < x; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
