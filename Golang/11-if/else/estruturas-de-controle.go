package main

import "fmt"

func main() {
	numero := 10

	//Não é obrigatório o uso de () no entanto o uso de {} é obrigatório
	if numero > 10 {
		fmt.Println("é maior que 15")
	} else {
		fmt.Println("é menor ou igual a 15")
	}

	//if init
	//A variável declarada na condição do if só pode ser usada dentro do escopo do if
	//logo se tentar acessá-la após o if o compilador irá apontar um erro
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	}
}
