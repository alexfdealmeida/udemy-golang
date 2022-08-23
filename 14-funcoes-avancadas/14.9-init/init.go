package main

import "fmt"

var n int

// A funcao init eh executada antes da main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada!")
	fmt.Println(n)
}
