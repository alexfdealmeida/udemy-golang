package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Alex"
	u.idade = 37
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Leidiane", 33, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Teste"}
	fmt.Println(u3)
}
