package main

import "fmt"

func notaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n < 3 && n >= 0:
		return "E"
	default:
		return "Invalid note"
	}
}

func main() {
	fmt.Println(notaConceito(9.8))
	fmt.Println(notaConceito(5.6))
	fmt.Println(notaConceito(2))
	fmt.Println(notaConceito(-2))
}
