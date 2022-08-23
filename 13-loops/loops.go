package main

import (
	"fmt"
)

func main() {
	//i := 0

	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}

	//fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando j", j)
	//	time.Sleep(time.Second)
	//}

	nomes := []string{"Alex", "Leide", "Teste"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		// No caso da letra, por se tratar de um char, vai retornar a posicao na tabela ASCII
		// Para exibir a letra correspondente, deve-se utilizar a funcao string
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Alex",
		"sobrenome": "Almeida",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//for {
	//	fmt.Println("Loop infinito")
	//	time.Sleep(time.Second)
	//}
}
