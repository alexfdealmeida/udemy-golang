package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// Desta forma, vai perder tempo, pois vai aguardar 2 segundos para receber a mensagem do canal2... ou seja, vai imprimir: Canal 1, Canal 2, Canal 1, Canal 2...
		//mensagemCanal1 := <-canal1
		//fmt.Println(mensagemCanal1)
		//mensagemCanal2 := <-canal2
		//fmt.Println(mensagemCanal2)

		// Assim vai imprimir: Canal 1, Canal 1, Canal 1, Canal 1, Canal 2...
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
