package main

import "fmt"

func main() {
	fmt.Print("Mesma linha ")
	fmt.Print("na mesma linha\n")

	fmt.Println("Nova linha")
	fmt.Println("Outra nova linha")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa!"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, c, d, x)
}
