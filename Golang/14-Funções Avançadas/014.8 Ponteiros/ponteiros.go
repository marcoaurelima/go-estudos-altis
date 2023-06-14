package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

//Função esperando um endereço de memória como parâmetro
func inverterSinalComPonteiro(num *int) {
	//Usando * para pegar o valor do endereço de memória
	*num *= -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	//Passando o endereço de memória como parâmetro usando &
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
