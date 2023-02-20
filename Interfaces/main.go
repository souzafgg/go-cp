package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func area(f forma) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

func main() {
	r := retangulo{10, 15}
	area(r)

}
