package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println("Aritmeticos:")
	fmt.Println(soma)
	fmt.Println(subtracao)
	fmt.Println(divisao)
	fmt.Println(multiplicacao)
	fmt.Println(restoDivisao)

	// Atribuicao
	var variavel1 string = "String"
	variavel2 := "String 2"

	fmt.Println("Atribuicao:")
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println("Relacionais:")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Logicos
	verdadeiro, falso := true, false
	fmt.Println("Logicos:")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// Unarios
	fmt.Println("Unarios:")
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 15
	numero *= 2
	numero /= 2
	fmt.Println(numero)

	// Ternario
	// nao tem suporte: numero > 5 ? "Maior que 5" : "Menor que 5"
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println("Ternario:")
	fmt.Println(texto)
}
