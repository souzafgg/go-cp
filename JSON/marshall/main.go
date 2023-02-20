package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {
	c := dog{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	c2 := map[string]string{
		"nome": "Hana",
		"raça": "Shih",
	}

	c2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2JSON)
	fmt.Println(bytes.NewBuffer(c2JSON))

	c3 := map[string]string{
		"nome": "Hana",
		"raça": "Shih",
	}
	c3JSON, _ := json.Marshal(c3)

	fmt.Println(c3JSON)
	fmt.Println(bytes.NewBuffer(c3JSON))

}
