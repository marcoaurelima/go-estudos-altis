package main

import "fmt"

func main() {
	fmt.Println("Executando a função main")
}

//A função init é executanda antes da função main
func init() {
	fmt.Println("Executando a função init")
}
