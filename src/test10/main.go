package main

import "fmt"

func main() {
	a := make([]int, 1)
	a = append(a, 1, 2, 3)

	fmt.Printf("切片的长度为%d 容量为%d\n", len(a), cap(a))

	people := []string{"Sam", "Bom", "Jack"}
	people1 := make([]string, 3)
	var people2 []string
	people3 := make([]string, 0, 5)

	fmt.Println(people, people1, people2, people3)

}
