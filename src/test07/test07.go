package main

import (
	"fmt"
	"reflect"
)

func main()  {
	str1 := "Goland"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[7]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c \n", str2[2], str2[2])
	fmt.Println("len(str2)", len(str2))
	fmt.Println(str1[0])
	fmt.Println(str1[1])
	fmt.Println(str1[2])
	fmt.Println(str1[3])
	fmt.Println(str1[4])
	fmt.Println(str1[5])
	x := 1
	fmt.Println(x)
	{
		x = 2
		fmt.Println(x)
		x := 3
		fmt.Println(x)
	}
	fmt.Println(x)
}