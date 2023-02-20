package main

import "fmt"

func func1() {
	fmt.Println("Executando a função 1")
}

func func2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("O resultado será retornado")
	fmt.Println("Entrando função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}
	return false
}

func main() {

	// defer func1()
	// // defer = adiar a função até o final (nesse caso vai executar a função2 e depois a função1)
	// func2()

	fmt.Println(alunoAprovado(5, 5))

}
