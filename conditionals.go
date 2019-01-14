package main

import "fmt"

func main() {
	ifTest()
}

func ifTest() {
	var number int
	fmt.Println("Ingresa un numero")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	if number := 2; number == 2 {
		fmt.Println("entro al condicional")
	}
}
