package main

import "fmt"

func main() {

	var number int
	fmt.Println("Digite un numero")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("El numero es 1")

	default:
		fmt.Println("El numero no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("El numero es par")
	default:
		fmt.Println("El numero es impar")
	}
}
