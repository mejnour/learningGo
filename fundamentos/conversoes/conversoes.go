package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println((x) / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...concatenacao de str com cast de num -> str busca na tabela unicode
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste" + strconv.Itoa(123))

	// str para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// str to bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	b1, _ := strconv.ParseBool("1")
	if b1 {
		fmt.Println("Verdadeiro")
	}
}
