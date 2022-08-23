package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println("String", 1, true)

	// Nao recomendavel
	mapa := map[interface{}]interface{}{
		1:        "string",
		"string": 1,
	}

	fmt.Println(mapa)
}
