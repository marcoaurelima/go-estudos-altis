package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint
	altura uint
}

// ATENÇÃO: não deve ser uma propriedade chamada pessoa do tipo pessoa
// É apenas pessoa, sem precisar passar o tipo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa := pessoa{"Alan", 23, 175}
	estudante := estudante{pessoa, "Software", "UFC"}

	fmt.Println(estudante)
}
