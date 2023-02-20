package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{
		nome:   "Adalberto",
		idade:  20,
		altura: 184,
	}

	p2 := estudante{
		pessoa: pessoa{
			nome:   "Hana",
			idade:  8,
			altura: 30,
		},
		curso:     "Redes",
		faculdade: "Infnet",
	}

	fmt.Println(p1, p2)

}
