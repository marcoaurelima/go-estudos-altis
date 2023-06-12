package main

import "fmt"

func main() {
	//Declarando canal com buffer, ou seja ele tem um tamanho máximo de 2
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Olá mundo 2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
