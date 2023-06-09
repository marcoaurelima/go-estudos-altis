package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Posicao 0", "Posicao 1"}
	fmt.Println(array2)

	//Aqui irá definir o tamanho do array de acordo com o número de itens informados entre as {}
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	//Para definir um slice basta não definir o tamanho máximo do array junto com as {} que podem
	//possuir ou não valores
	//Slice é basicamente um array sem um tamanho pré definido, ou seja, possui tamanho dinâmico
	slice := []int{}
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	//Usa-se append para adicionar um novo valor a um array
	slice2 = append(slice2, 6)
	fmt.Println(slice2)

	fmt.Println("-----------------------")
	fmt.Println("Arrays Internos")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Length - tamanho
	fmt.Println(cap(slice3)) //Capacity - capacidade
}
