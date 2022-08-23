package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Executando tarefa 1", canal)

	fmt.Println("Início do programa")

	// Neste caso, resultaria num deadlock, pois o canal permaneceria aberto e não teria mais dados pra receber.
	// Deve-se ter atencao, pois trata-se de um erro que ocorre somente em tempo de execucao.
	//for {
	//	mensagem := <-canal
	//	fmt.Println(mensagem)
	//}

	//Implementacao 1
	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}

	//Implementacao 2
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
