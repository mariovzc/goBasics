package main

import (
	"fmt"

	"./name"
)

func main() {
	stringUTF8 := name.GetUnicode()
	hello := "hola"
	fmt.Println("Cadena con UTF-8: ", stringUTF8)
	fmt.Println(hello[0])
	fmt.Println(string(hello[0]))
	fmt.Println(len(hello))
}
