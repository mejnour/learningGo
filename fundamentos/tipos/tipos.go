package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 310000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	test := math.MaxUint8
	fmt.Print(test)

	// sem sinais -- uint8, uint16, uint32, uint64
	var b byte = 255   // uint8
	var b1 uint8 = 255 // uint8
	fmt.Println("O valor de b é", b)
	fmt.Println("O byte é", reflect.TypeOf(b))
	fmt.Println("O uint8 é", reflect.TypeOf(b1))

	var b2 uint16 = 65535 // uint16
	fmt.Println("O valor de b2 é", b2)
	fmt.Println("O uint16 é", reflect.TypeOf(b2))

	var b3 uint32 = 4294967295 // uint32
	fmt.Println("O valor de b3 é", b3)
	fmt.Println("O uint32 é", reflect.TypeOf(b3))

	var b4 uint64 = 18446744073709551615 // uint64
	fmt.Println("O valor de b4 é", b4)
	fmt.Println("O uint64 é", reflect.TypeOf(b4))

	var c rune = 'a' // rune é um alias tabela unicode para int32
	fmt.Println("O valor de c é", c)
	fmt.Println("O rune é", reflect.TypeOf(c))

	// numeros reais
	var x float32 = 49.99 // float32
	fmt.Println("O valor de x é", x)
	fmt.Println("O float32 é", reflect.TypeOf(x))

	var y float64 = 49.99 // float64
	fmt.Println("O valor de y é", y)
	fmt.Println("O float64 é", reflect.TypeOf(y))

	// boolean
	bo := true
	fmt.Println("O valor de bo é", bo)
	fmt.Println("O boolean é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	bo2 := 1
	fmt.Println("O valor de bo2 é", bo2)
	fmt.Println("O tipo de bo2 é", reflect.TypeOf(bo2))
	// fmt.Println(!bo2) // bo2 é um inteiro, não um booleano, então o operador ! não é válido aqui

	// string
	s1 := "Olá, mundo!"
	fmt.Println("O valor de s1 é", s1)
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println(s1 + "!!!")
	fmt.Println("O tamanho de s1 é", len(s1))

	s2 := `Olá, 
	mundo!`
	fmt.Println("O valor de s2 é", s2)

	// char
	char := 'a'
	// var char2 char = 'b' // char não é um tipo válido
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}


(50*3)+(11,63*3)+(27,90*3)+(29,12*3)+(12,48*3)+(29,90*3)+(200*3) = 1.083,09
95 + 31,96 + 31,79 = 158,75
(300*3)+(162*3) = 1.386
1.386 - 158,75 - 1.083,09 = 144,16