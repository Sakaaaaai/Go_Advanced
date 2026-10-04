//go:build ignore

package main

import "fmt"

func createOnStack() int {
	x := 10
	return x
}

func createOnHeap() *int {
	y := 20
	return &y
}

func escapeByInterface() {
	z := 30
	fmt.Println(z)
}

func main() {
	val := createOnStack()
	ptr := createOnHeap()

	fmt.Printf("Stack value: %d\n", val)
	fmt.Printf("Heap value: %d\n", *ptr)
	escapeByInterface()
}
