package main

import "fmt"

func main() {

	u := map[string]int{
		"nome":       1,
		"sobrenome":  2,
		"quantidade": 5,
	}

	fmt.Println(u)
	fmt.Println(u["quantidade"])

	u2 := map[string]map[string]int{
		"amount": {
			"name":      5,
			"sobrenome": 25,
		},
		"another": {
			"jjj": 1,
			"kkk": 2,
		},
	}
	fmt.Println(u2)
	delete(u2, "amount")

	u2["signo"] = map[string]int{
		"nome": 2,
	}
}
