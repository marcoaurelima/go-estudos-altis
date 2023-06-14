package main

import "fmt"

//Função de soma recebendo dois inteiros e retornando um inteiro
func somar(num1 int, num2 int) int {
	return num1 + num2
}

//Função retornando multiplos valores
//Lembrando que podem ser valores de tipos diferentes Ex: int e string
func calculos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(1, 3)
	fmt.Println(soma)

	//func é um tipo, logo pode ser atribuidas a variáveis
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	retorno := f("teste")
	fmt.Println(retorno)

	resultadoSoma, resultadoSub := calculos(1, 3)
	fmt.Println(resultadoSoma, resultadoSub)

	//Caso não queira o retuno de uma função que retorna multiplos valores
	//basta colocar um _
	resultadoSoma2, _ := calculos(1, 3)
	fmt.Println(resultadoSoma2)
}
