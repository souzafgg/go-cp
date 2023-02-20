package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func main() {

	user1 := usuario{
		"Usuário 1", 20,
	}
	user1.salvar()

	user2 := usuario{"Adalberto", 20}
	user2.salvar()

	c := user2.maiorIdade()
	fmt.Println(c)
}
