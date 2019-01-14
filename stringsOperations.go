package main

import (
	"fmt"
	"strings"
)

func main() {

	var text = "Hello World, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))

	fmt.Println(strings.ToLower(text))

	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))

	fmt.Println(strings.Split(text, ","))
}
