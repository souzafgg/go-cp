package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//paralelismo != concorrência
	//invocar goroutine usar go

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Olá, mundo")
		waitGroup.Done()
	}()
	go func() {
		escrever("Programando em go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
