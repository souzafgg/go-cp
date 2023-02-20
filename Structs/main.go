package main

import (
	"fmt"
)

type user struct {
	nome     string
	idade    uint8
	endereco string
}

type endereco struct {
	rua    string
	numero int8
}

func main() {

	var u user
	u.nome = "Adal"
	u.idade = 20
	fmt.Println(u)


	u2 := user{
		endereco: endereco{
			nome: "Adal",
			idade: 20,
		
		endereco{
			rua: "Rua",
			numero: 5,
			}
		}
	}


	fmt.Println(u2)

}
