package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador que vai contar
	fmt.Println(numeros)

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}

	// var teste int = 6
	// for _, num := range teste {
	// 	fmt.Println(num)
	// }
}
