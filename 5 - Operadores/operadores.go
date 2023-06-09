package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 1 + 3
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 1 * 2
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	//ATRIBUIÇÃO
	var texto string = "texto da string"
	variavelTexto := "texto da segunda string"
	fmt.Println(texto, variavelTexto)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 3)
	fmt.Println(1 >= 3)
	fmt.Println(1 < 3)
	fmt.Println(1 <= 3)
	fmt.Println(1 == 3)
	fmt.Println(1 != 3)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 15
	numero *= 3
	numero /= 3
	numero %= 3

	fmt.Println(numero)

	//OPERADOR TERNÁRIO

}
