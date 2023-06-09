package main

import "fmt"

func main(){
 
  // Declaração de tipos compostos devem ser feitos através do make()
  s := make([]int, 5) // Slice de inteiros de tamanho 5, com valor inicial 0
  m := make(map[string]string) // Map de chave string e valor string
  c := make(chan string) // Canal de comunicação entre gorotines

  fmt.Println(s, m, c)

}

