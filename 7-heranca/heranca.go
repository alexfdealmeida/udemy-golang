package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Alex", 37, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Sistema de Informacao", "UEG"}
	fmt.Println(e1.nome, e1.idade, e1.altura, e1.curso, e1.faculdade)
}
