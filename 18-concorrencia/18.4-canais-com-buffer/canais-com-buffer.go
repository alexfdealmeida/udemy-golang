package main

import "fmt"

func main() {
	// O canal com buffer, so vai bloquear a aplicacao quando ele atingir a capacidade maxima (buffer)
	canal := make(chan string, 2)
	canal <- "Hello World!"
	canal <- "Go Lang"

	// Caso receba um novo valor, resultara em deadlock, pois vai exceder a capacidade (2)
	//canal <- "New Value"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
