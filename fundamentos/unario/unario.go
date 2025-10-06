package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(x, y)

	// postfix
	x++
	fmt.Println(x, y)

	y--
	fmt.Println(x, y)
}
