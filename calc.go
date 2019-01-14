package main

import "fmt"

func main() {
	var num1, num2 float32
	fmt.Println("Bienvenido !")
	fmt.Println("Digite el primer numero")
	fmt.Scanf("%f", &num1)

	fmt.Println("Digite el segundo numero")
	fmt.Scanf("%f \n", &num2)

	sum, res := getValues(num1, num2)

	fmt.Printf("%f %f", sum, res)
}

func getValues(n1 float32, n2 float32) (float32, float32) {
	return sum(float32(n1), float32(n2)), res(float32(n1), float32(n2))
}

func sum(n1 float32, n2 float32) float32 {
	return n1 + n2
}

func res(n1 float32, n2 float32) float32 {
	return n1 - n2
}
