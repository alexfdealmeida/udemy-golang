package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"

	variavel2 := "Variável 2"

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	const constante1 string = "Constante 1"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	fmt.Println(constante1)

	// Inverter valores das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
