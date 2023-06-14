package main

import "fmt"

func main() {
	// Operações de enviar e receber dados via channels são BLOQUEANTES
	// O segundo envio "Olá mundo 2" fez o canal chegar no limite (2),
	// o que o fez liberar o fluxo de execução da função main.
	// Algo como canal <- "Olá mundo 3" resultaria em deadlock

	//Declarando canal com buffer, ou seja ele tem um tamanho máximo de 2
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Olá mundo 2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
