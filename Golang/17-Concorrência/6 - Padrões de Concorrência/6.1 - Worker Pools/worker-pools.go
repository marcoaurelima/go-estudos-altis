package main

import "fmt"

// este padrão é ideal quando queremos realizar várias tarefas em paralelo,
// sendo estas tarefas INDEPENDENTES entre si.
// Por exemplo, um programa que gerencie vários pedidos de uma loja online.

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Workers (trabalhadores) que irão executar tudo.
	// Idealmente, a quantidade de workers deve ser igual a quantidade
	// de núcleos do procesdsador, para que cada worker fique com 1 core.
	// Repare que são 4 workers diferentes, mas estão compartilhando o mesmos canais
	// Durante o primeiro for, serão entregues mensagens e o "trabalhador" que estiver
	// disponível no momento executará ele.
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

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

// funcão tem como parâmetro um canal de somente de entrada e outro somente de saída
// ele vai percorrer todas as tarefas.
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
