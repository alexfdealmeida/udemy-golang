package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO

	go escrever("Executando tarefa 1")
	escrever("Executando tarefa 2")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
