package main

import (
	"fmt"
	"time"
)

func main() {
	c := escrever("Olá, mundo")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func escrever(t string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", t)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
