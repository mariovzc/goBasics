package main

import "fmt"

const helloWorld = "Hola %s %s, bienvenido al mundo de GO. \n"

func main() {

	// solo se utiliza para inizializar variables
	lastname := "Vzc"

	//var number = 100
	number := sum(50, 50)
	//var (
	//	a = 10
	//	b = 20
	//	c = 30
	//)

	a, b, c := getVariables()

	f32, f64 := getFloat()

	name := getName()

	fmt.Printf(helloWorld, name, lastname)

	fmt.Println(a, b, c, number)

	fmt.Println(f32)
	fmt.Println(f64)
}

func getName() string {
	var name string
	name = "sin nombre"

	fmt.Println("Ingresa tu nombre")
	fmt.Scanf("%s", &name)

	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

// float 64 tiene mejor precision
func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
