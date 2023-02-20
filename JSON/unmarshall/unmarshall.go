package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {
	cJSON := `{"nome":"Hana","raça":"Shih"}`

	var c dog

	json.Unmarshal([]byte(cJSON), &c)

	fmt.Println(c)
}
