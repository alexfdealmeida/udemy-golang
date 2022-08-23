package main

import "fmt"

func main() {
	fmt.Println("Atribuicao por valor")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	fmt.Println("Atribuicao por referencia (ponteiro)")

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	// Posicao de memoria
	fmt.Println(variavel3, ponteiro)
	// Valor na posicao de memoria
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
