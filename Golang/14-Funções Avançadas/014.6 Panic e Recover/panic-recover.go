package main

import "fmt"

func recuperarExecucao() {
	//O recover irá recuperar a execução de uma clásula panic
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//O panic mata a execução
	//No entanto antes de para a execução o panic executa as funções com "defer"
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução")
}
