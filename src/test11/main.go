package main

import "fmt"

func ssa(a, b, c uint) uint {
	return a * b * c
}

func main() {
	fmt.Println(ssa(1, 2, 3))
}


