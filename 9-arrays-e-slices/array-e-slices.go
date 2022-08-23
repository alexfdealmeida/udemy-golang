package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays:")

	var array1 [5]string
	array1[0] = "P1"
	fmt.Println(array1)

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array2)

	// Definir o tamanho de acordo com a quantidade de elementos
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println("Slices:")

	//slice = array dinamico
	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	slice = append(slice, 15)
	fmt.Println(slice)

	// indice 1 = inclusivo
	// indice 3 = exclusivo
	// Nao se trata de uma copia, e sim de um ponteiro...
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// ... portanto, se alterar o array, ira refletir no slice
	array2[1] = "P2 alterada"
	fmt.Println(slice2)

	fmt.Println("Arrays internos:")

	// tipo, tamanho inicial, capacidade (opcional)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	// tamanho atual do slice
	fmt.Println(len(slice3))
	// capacidade do slice
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	// Ao exceder a capacidade, adicionando 12 itens, ela sera dobrada automaticamente (24)
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// Quando a capacidade nao for informada, ira corresponder ao tamanho do slice
	slice4 := make([]float64, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	//fmt.Println(reflect.TypeOf(slice))
	//fmt.Println(reflect.TypeOf(array3))
}
