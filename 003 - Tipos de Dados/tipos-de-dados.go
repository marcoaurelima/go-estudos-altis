package main

import (
	"errors"
	"fmt"
)

func main() {
	//Positivos e negativos
	var varInt int = 1
	var varInt8 int8 = 1
	var varInt16 int16 = 1
	var varInt32 int32 = 1
	var varInt64 int64 = 1

	fmt.Println(varInt, varInt8, varInt16, varInt32, varInt64)

	//Inteiro sem sinais + ou -
	var varUint uint = 1
	var varUint8 uint8 = 1
	var varUint16 uint16 = 1
	var varUint32 uint32 = 1
	var varUint64 uint64 = 1

	fmt.Println(varUint, varUint8, varUint16, varUint32, varUint64)

	//INT32 = rune
	var varRune rune = 1
	fmt.Println(varRune)

	//BYTE = uint8
	var varByte byte = 1
	fmt.Println(varByte)

	//Float
	var varFloat float32 = 0.1
	fmt.Println(varFloat)

	//Float
	var varFloat64 float64 = 0.1
	fmt.Println(varFloat64)

	//String, deve sempre ser entre aspas duplas
	var varString string = "teste"
	fmt.Println(varString)

	//Pegar número do caractere na tabela ascii,
	//só pode passar um caractere e deve ser entre aspas simples
	char := 'B'
	fmt.Println(char)

	//valor zero: é o valor que é atribuida a variável caso ela não seja incializada
	//valor zero de string é ""
	var stringVazia string
	fmt.Println(stringVazia)

	//Tipo booleano
	//Valor padrão é false
	var varBool bool = true
	fmt.Println(varBool)

	//Tipo Erro
	//valor padrão é nil
	var varErro error
	varErro = errors.New("Texto do seu erro")
	fmt.Println(varErro)

	// Tipo interface (genérico)
	// Usado para valores dinâmicos que não se sabe o tipo de antemão
	quadrado := func(a interface{})interface{} {
    switch v := a.(type) {
		case int:
			return v * v
		case float64:
			return v * v
		default:
			return nil
		}
	}
  
  fmt.Println(quadrado(3))
  fmt.Println(quadrado(2.5))

}
