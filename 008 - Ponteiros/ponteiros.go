package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++

	fmt.Println(variavel1, variavel2)
	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int = 100
	//Usando * para declarar um ponteiro
	var variavel4 *int
	//Usando & para passar o endereço da memória para o ponteiro
	variavel4 = &variavel3

	//Vai printar o valor da variavel3 e seu endereço na memória
	fmt.Println(variavel3, variavel4)

	//Usando * para desreferenciar o endereço da memória e retornar o valor
	//Vai printar duas vezes o valor da variável 3
	fmt.Println(variavel3, *variavel4)
}
