// Mapeando um JSON para uma STRUCT (deserialização)

package main

import (
	"bytes"
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
	Idade     int    `json:"idade"`
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

	//Converter struct para json
	//variável que irá ser convertida
	person := Pessoa{
		"Alan",
		"Nogueira",
		23,
	}

	//Converter map para json
	//variável que irá ser convertida
	newPerson := map[string]interface{}{
		"nome":      "Alan",
		"sobrenome": "nogueira",
		"idade":     23,
	}

	//A função Marshal retorna um slice de bytes(referente ao JSON) e um erro
	//convertendo struct
	personEmJSON, erro := json.Marshal(person)
	//Tratando o erro caso ocorra
	if erro != nil {
		log.Fatal(erro)
	}
	//convertendo map
	newPersonEmJSON, erro := json.Marshal(newPerson)
	//Tratando o erro caso ocorra
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("---------struct---------")
	//Printando o json em bytes
	fmt.Println(personEmJSON)
	//Convertendo o slice de bytes em json
	fmt.Println(bytes.NewBuffer(personEmJSON))
	fmt.Println("---------fim---------")
	fmt.Println("---------map---------")
	//Printando o json em bytes
	fmt.Println(newPersonEmJSON)
	//Convertendo o slice de bytes em json
	fmt.Println(bytes.NewBuffer(newPersonEmJSON))
	fmt.Println("---------fim---------")

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
