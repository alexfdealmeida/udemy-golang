package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(10)

	for i := uint(1); i <= posicao; i++ {
		fmt.Print(fibonacci(i), " ")
	}

	//fmt.Println(fibonacci(posicao))
}
