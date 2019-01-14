package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println(c)
	}

	s := 1000

	for {
		s--

		if s == 0 {
			fmt.Println("Terminado")
			break
		}
	}
}
