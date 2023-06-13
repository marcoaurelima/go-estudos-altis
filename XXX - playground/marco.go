package main

import "fmt"

// este padrão é ideal quando queremos realizar várias tarefas em paralelo,
// sendo estas tarefas INDEPENDENTES entre si.
// Por exemplo, um programa que gerencie vários pedidos de uma loja online.

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Workers (rabalhadores) que irão executar tudo.
	// Idealmente, a quantidade de workers deve ser igual a quantidade
	// de núcleos do procesdsador, para que cada worker fique com 1 core.
	go worker("worker1", tarefas, resultados)
	go worker("worker2", tarefas, resultados)
	go worker("worker3", tarefas, resultados)
	go worker("worker4", tarefas, resultados)

	// Laço for envia 45 mensagens para o canal, que é sua capacidade máxima.
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	// Avisa que todas as mensagens foram enviadas e fecha o canal fica somente leitura)
	// Isso também habilita o uso do for mais a frente, pois terminda quais quer processamentos.
	close(tarefas)

	// Laço for captura os 45 resultados pelo canal de saída e em seguida printa
	// tudo na tela.
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(workerName string, tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		fmt.Println(workerName, numero)
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
