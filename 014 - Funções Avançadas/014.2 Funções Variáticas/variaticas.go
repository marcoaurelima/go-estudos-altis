package main

import "fmt"

//Funções variáticas são funções que podem receber um numero indeterminado de parâmetros
//para definir isso se colocar o nome da varável seguido de ... e o tipo
//(nomeDaVariavel ...int)
//ao entrar na função virá um slice de tipo inteiro
//Por exemplo caso chame a função abaixo passando os valores (1, 2, 3, 4, 5, 6, 7)
//A variável numeros será um slice com o seguinte valor [1, 2, 3, 4, 5, 6, 7]
func soma(numeros ...int) (soma int) {
	for _, num := range numeros {
		soma += num
	}
	return
}

//É possível passar outros parâmetros para uma função variática, no entando,
//Deve haver apenas um parâmetro de quantidade indeterminada e ele deve ser o último
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(soma())
	escrever("Texto", 1, 2, 3, 4, 5, 6, 7)
}
