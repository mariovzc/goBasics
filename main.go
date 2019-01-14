package main

import (
	"fmt"

	"./maps"
	//"./name"
)

const helloWorld = "Hola %s %s, bienvenido al mundo de GO. \n"

func main() {

	/*
		// solo se utiliza para inizializar variables
		lastname := "Vzc"

		//var number = 100
		//var (
		//	a = 10
		//	b = 20
		//	c = 30
		//)

		a, b, c := getVariables()

		firstName := name.GetName()

		f32, f64 := getFloat()

		fmt.Printf(helloWorld, firstName, lastname)

		fmt.Println(a, b, c)

		fmt.Println(f32)
		fmt.Println(f64)*/

	fmt.Println(maps.GetMap("yohan"))

	nombre := "platzi"
	fmt.Println("Hola %s", nombre)
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

// float 64 tiene mejor precision
func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
