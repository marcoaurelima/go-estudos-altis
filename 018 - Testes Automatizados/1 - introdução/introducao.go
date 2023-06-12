package main

import (
	"fmt"
	enderecos "introducao_testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida")
	fmt.Println(tipoEndereco)
}
