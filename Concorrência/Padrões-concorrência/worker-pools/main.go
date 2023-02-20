package main

import "fmt"

func main() {

	// fibonacci 1 1 2 3 5 8 13
	t := make(chan int, 45)
	r := make(chan int, 45)

	go worker(t, r)
	for i := 0; i < 0; i++ {
		t <- i
	}
	close(t)

	for i := 0; i < 45; i++ {
		x := <-r
		fmt.Println(x)
	}
}

func worker(t <-chan int, r chan<- int) {
	for numero := range t {
		r <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
