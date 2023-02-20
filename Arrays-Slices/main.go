package main

import "fmt"

func main() {

	var array [5]int
	array[4] = 1
	fmt.Println(array)

	array2 := []string{"P1", "P2"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(array3)

	//slice

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice)

	//arrays internos

	slice3 := make([]float64, 5, 10)
	slice3 = append(slice3, 1)

	fmt.Println(slice3, cap(slice3), len(slice3))
}
