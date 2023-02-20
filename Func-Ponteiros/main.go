package main

import "fmt"

func invSinal(n int) int {
	return n * -1
}

func InvSinalPointer(n *int) {
	*n = *n * -1
}

func main() {

	numero := 20
	nInvertido := invSinal(numero)
	fmt.Println(nInvertido)

	novoN := 40
	InvSinalPointer(&novoN)
	fmt.Println(novoN)

}
