package main

import (
	"fmt"
	"time"
)

func main() {
	// Encapsular chamada de um goroutine, retornando em um canal
	canal := escrever("Hello World!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
