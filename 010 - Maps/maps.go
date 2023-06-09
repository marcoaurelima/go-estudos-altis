package main

import "fmt"

func main() {
	//Dentro dos [] fica os tipos das chaves
	//E fora fica o tipo dos valores
	//Não é possível ter chaves de tipos diferentes ou valores de tipos diferentes
	//Ex: usuario := map[int]string{}
	usuario := map[string]string{
		"nome":      "Alan",
		"sobrenome": "Nogueira",
	}
	fmt.Println(usuario)

  
  // make para chaves de tipos diferentes
  pessoa := make(map[interface{}]string)
  pessoa[0] = "31"
  pessoa["nome"] = "marco"
  pessoa[3.14] = "PI"
  
  fmt.Printf("%s | %s | %s\n", pessoa["nome"], pessoa[0], pessoa[3.14])

	//Para acessar o valor usa-se nomeDaVariavel["chave"]
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Alan",
			"ultimo":   "Nogueira",
		},
		"curso": {
			"nome":   "ES",
			"cidade": "Russas",
		},
	}
	fmt.Println(usuario2)

	//Para remover um valor de um map usa-se o método delete
	//Passando como parâmetro o map e a chave
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	//Para adicionar basta fazer um atribuição de valor apartir da chave
	usuario2["endereco"] = map[string]string{
		"rua": "Miguel Pereira",
	}
}
