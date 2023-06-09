package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função1")
}

func funcao2() {
	fmt.Println("Executando função2")
}

func alunoAprovado(n1, n2 int) bool {
	//O defer faz que o print ocorra somente logo antes do return
	defer fmt.Println("Média calculada")
	fmt.Println("Entrando na função para verificar aprovação")

	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	}

	return false
}

func main() {
	//defer adia a execução
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoAprovado(7, 8))
}
