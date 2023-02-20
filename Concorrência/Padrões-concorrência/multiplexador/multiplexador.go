package main

import (
	"fmt"
	"time"
)

func main() {
	canal := mx(escrever("Olá, mundo"), escrever("Parte2"))
	for i:=0; i<10; i++{
		fmt.Println(<-canal)
	}
}

func mx(canal1, canal2 <-chan string) <-chan string
	canalSaída := make(chan string)
go func() {  
	for {
		select {
		case mensagem := <- canal1:
			canalSaída <- mensagem
		case mensagem := <- canal2:
			canalSaída <- mensagem
		}s
	}
}()
return canalSaída

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
