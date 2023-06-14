package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Ok Go")

	erro := checkmail.ValidateFormat("alan.nogueira91@gmail.com")
	fmt.Println(erro)
}
