package main

import (
	"fmt"
	"time"
)

func main() {
	//declara o canal
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//Envia o valor pelo canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//Fecha o canal
	close(canal)
}
