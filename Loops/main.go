package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 5 {
		fmt.Println("....")
		time.Sleep(time.Second)
		i++
	}
	fmt.Println(i)

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	n := [3]string{"João", "Davi", "Lucas"}

	for q, r := range n {
		fmt.Println(q, r)
	}

	for o, p := range "PALAVRA" {
		fmt.Println(o, string(p))
	}

	// for com range em map | Obs: Não há como fazer em struct
	u := map[string]int{
		"Nome":  2,
		"Idade": 20,
	}

	for c, v := range u {
		fmt.Println(c, v)
	}

}
