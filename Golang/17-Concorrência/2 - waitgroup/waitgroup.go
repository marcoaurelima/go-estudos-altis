package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Teste")
		waitGroup.Done()
	}()

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Indica que deve-se aguardar o fim das goroutines
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
