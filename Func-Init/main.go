package main

import "fmt"

// função init será sempre iniciada em primeiro, independente de onde esteja no código
func init() {
	fmt.Println("Iniciando a função init")
}

func main() {

	fmt.Println("Print na main da main")
}
