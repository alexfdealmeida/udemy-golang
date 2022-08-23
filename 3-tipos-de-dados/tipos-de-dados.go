package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32 ou int64
	// int ira definir o tamanho de acordo com a arquitetura do computador (32 ou 64)
	var numero int = 100
	fmt.Println(numero)

	// uint aceita somente valores positivos
	// caso seja informado um valor negativo, resultara em erro: cannot use -1000 (untyped int constant) as uint32 value in variable declaration (overflows)
	//var numero2 uint32 = -1000
	//fmt.Println(numero2)

	// alias
	// rune = int32
	// byte = int8
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	// float32 ou float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.45
	fmt.Println(numeroReal2)

	// utilizando a inferencia, sera definido o tamanho de acordo com a arquitetura do computador (32 ou 64)
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Sera exibido o numero do caracter na tabela ASCII: 66
	char := 'B'
	fmt.Println(char)

	// Valores default
	// string = ""
	// numeric = 0
	// bool = false
	// error = nil
	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
