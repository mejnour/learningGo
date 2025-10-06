package main

import "fmt"

func main() {
	i := 1

	// Go nao tem aritmetica de ponteiros
	var p *int = nil
	fmt.Println(p)
	p = &i          // ponteiro para endereço de i
	fmt.Println(p)  // printa endereço de i
	fmt.Println(*p) // printa valor do endereço de i
	*p++
	fmt.Println(*p) // printa valor desreferenciado e incrementado do endereço de i

	i++
	fmt.Println(i)
	fmt.Println(*p)
}
