package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for com range
	nomes := [3]string{"nome1", "nome2", "nome3"}
	for indice, nome := range nomes {
		fmt.Println(nome, indice)
	}
	//ou
	//Assim não é forçado o uso da variável indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//Iterando string
	for _, letra := range "PALAVRA" {
		//Valor da tabela ascii
		fmt.Println(letra)
		//Letra
		fmt.Println(string(letra))
	}

	//iterando map
	usuario := map[string]string{
		"nome":      "Alan",
		"sobrenome": "Nogueira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
