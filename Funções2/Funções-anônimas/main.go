package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Olá, mundo")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá, mundo")

	fmt.Println(retorno)

}
