package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

//Método atrelado a struct usuario
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário: %s\n", u.nome)
}

//O uso de ponteiros torna possível alterar os valores contidos na instância da struct
func (u *usuario) alterarNome(novoNome string) {
	u.nome = novoNome
}

func main() {
	usuario1 := usuario{"usuário 1", 20}
	fmt.Println(usuario1)
	//usuario1.salvar()
	usuario1.alterarNome("Novo nome do usuário")
	fmt.Println(usuario1)
}
