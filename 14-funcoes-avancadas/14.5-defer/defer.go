package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunEstaAprovado(n1, n2 float32) bool {
	// Ira adiar a execucao imediatamente antes do retorno da funcao
	defer fmt.Print("Resultado: ")

	fmt.Println("Verificando se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 5 {
		return true
	} else {
		return false
	}
}

func main() {
	//defer funcao1()
	//funcao2()

	fmt.Println(alunEstaAprovado(7, 8))
}
