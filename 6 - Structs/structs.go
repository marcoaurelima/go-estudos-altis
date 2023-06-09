package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint
}

func main() {

	fmt.Println("Structs em go")
	var user usuario
	user.nome = "Alan"
	user.idade = 23
	fmt.Println(user)

	user2 := usuario{"Alan", 23, endereco{"Miguel Pereira", 00}}
	fmt.Println(user2)

	user3 := usuario{nome: "Alan"}
	fmt.Println(user3)
}

//OBS Structs podem ser criadas e não usadas sem aparecer erro em desenvolvimento nem em execução
