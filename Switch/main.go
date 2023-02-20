package main

import "fmt"

func main() {

	dia := diaSemana(3)
	fmt.Println(dia)
}

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	default:
		return "Inválido"
	}

}
