package main

import "fmt"

func main() {
	stringUTF8 := getUnicode()
	hello := "hola"
	fmt.Println("Cadena con UTF-8: ", stringUTF8)
	fmt.Println(hello[0])
	fmt.Println(string(hello[0]))
	fmt.Println(len(hello))
}

func getUnicode() string {
	return "もしもし!"
}
