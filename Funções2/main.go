package main

import "fmt"

func calcMath(x, y int8) (soma int8, sub int8) {
	soma = x + y
	sub = x - y
	return
}

func main() {

	som, subt := calcMath(5, 4)
	fmt.Println(som, subt)

}
