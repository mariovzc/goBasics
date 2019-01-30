package main

import (
	"errors"
	"fmt"
)

//"./name"

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

	//fmt.Println(maps.GetMap("yohan"))

	//nombre := "platzi"
	//fmt.Println("Hola %s", nombre)
	//structs.InterfaceTest()
	//number, err := suma(2, 3)

	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(number)

	//pointerTest()

	//forGo(500)
	//forGo(400)

	// time.Sleep(10000 * time.Microsecond)
}

// go function
func suma(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor a es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}

	return a.(int) + b.(int), nil
}

// & trae la posicion en memoria
// * trae el valor del real del puntero
func pointerTest() {
	a := 100
	var b *int

	b = &a

	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)

	pointerModify(b)

	fmt.Println("Modificada")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c *int) {
	*c = 10
}

func helloGo(index int) {
	fmt.Println("Hola soy un print en la go routine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
		// allow to execute concurrency in go
	}
}
