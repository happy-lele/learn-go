package main

import "fmt"

func main() {
	// 类型转换
	var b byte = 100
	// var n int = b cannot use b (type byte) as type int in assignment
	var n int = int(b) // 显示转换
	println(n)

	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	fmt.Printf("%T\n", 'a')

	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac")

}

