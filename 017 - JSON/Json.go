// Mapeando um JSON para uma STRUCT (deserialização)

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Struct com a mesma estrutura do json
// Nome dos campos devem estar com letra maiúscula
// valor `json:""` é obrigatório, contendo o nome do campo do json que será mapeado
type Pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     int `json:"idade"`
}

func main() {

  // Declaração do json
	meuJson := `
[
  {
    "nome": "marco",
    "sobrenome": "aurelio",
    "idade": 31
  },
  {
    "nome": "carlos",
    "sobrenome": "silva",
    "idade": 22
  }
]`
  
  // Variável que vai receber a conversão do json
	var pessoa []Pessoa
  
  // função que faz a conversão JSON => STRUCT.
  // Primeiro parametro é o json, em formato de slice de byte (fiz um casting)
  // Segundo parametro é o end. de memória da struct que receberá o resultado
	err := json.Unmarshal([]byte(meuJson), &pessoa)
  
  // Gerenciar erros
	if err != nil {
		log.Println("erro: ", err)
	}
  
  // Printa o resultado
  fmt.Printf("Resultado:\n%v\n", pessoa)

}
