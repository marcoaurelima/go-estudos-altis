package main

import "fmt"

//Define-se o nome dos retornos
func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	//não é necessário colocar as variáveis no return
	return
}

func main() {
	soma, subtracao := calculos(5, 4)
	fmt.Println(soma, subtracao)
}
