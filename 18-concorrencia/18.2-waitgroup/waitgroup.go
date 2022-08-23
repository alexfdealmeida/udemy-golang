package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//NAO deixar a aplicacao finalizar enquanto a execução das rotinas nao finalizar
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Executando tarefa 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("Executando tarefa 2")
		waitGroup.Done()
	}()

	go func() {
		escrever("Executando tarefa 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("Executando tarefa 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
