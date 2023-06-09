package main

import "fmt"

func main() {

	func() {
		fmt.Println("Função anônima")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Função anônima com parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintln("Função anônima com retorno ->", texto)
	}("Parâmetro")

	fmt.Println(retorno)
}
