package main

import "fmt"

func main() {
	//Tipo explícito
	var variavel1 string = "variavel 1"

	//Inferência de tipo
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declarações multiplas
	var (
		variavel3 string = "asdasdfasdfdas"
		variavel4 string = "sadfasfasdfasd"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5: qfqwfqwefqwef", "variavel6: cqaecqwecwqec"

	fmt.Println(variavel5, variavel6)

	//constante
	const contante string = "constante"

	//trocando valor de duas variáveis
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)
}

/*
	Variáveis declaradas devem obrigatoriamente ser usadas
*/
