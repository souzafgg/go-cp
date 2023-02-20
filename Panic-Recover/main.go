package main

import "fmt"

func recExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6")

}

func main() {

	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Resultado pós execução")

}
