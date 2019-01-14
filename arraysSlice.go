package main

import "fmt"

func main() {
	getArray()
	getSlice()
}

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "array como texto"
	arr1[1] = "array2"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	slice1 = append(slice1, "2", "my", "slice")
	fmt.Println(slice1)
}
