package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Alex",
		"sobrenome": "Almeida",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	// Map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "ALex",
			"ultimo":   "Almeida",
		},
		"curso": {
			"nome":      "Sistemas de Informação",
			"faculdade": "UEG",
			"campus":    "Anápolis",
		},
	}

	fmt.Println(usuario2)

	fmt.Println(usuario2["nome"]["primeiro"])

	delete(usuario2, "curso")
	fmt.Println(usuario2)
}
