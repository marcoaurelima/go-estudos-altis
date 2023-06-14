package main

import (
	"fmt"
	"time"
)

func main() {
	//o goroutine é acionado ao colocar a palavra reservada go na frente de um método
	//Deste modo a linguagem diz que não precisa esperar o fim do método para seguir
	go escrever("Teste")
	escrever("Olá Mundo")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
