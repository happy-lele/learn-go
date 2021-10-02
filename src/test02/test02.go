package main

import "math"

const (
	a  byte = 100
	b  = 1e20
)

const (
	Sunday = iota 	// 0
	Monday 			// 1
	Tuesday  		// 2
	Wednesday  		// 3
	Thursday  		// 4
	Friday  		// 5
	Saturday 		// 6
)

const (
	_ = iota // iota = 0
	KB int64 = 1 << (10 * iota)
	MB      // 1 << (10 * 2)
	GB		// 1 << (10 * 3)
	TB		// 1 << (10 * 4)
)

const (
	A, B = iota, iota << 10	  // 0, 0 << 10
	C, D					 // 1, 1 << 10
)

const (
	A1 = iota   // 0
	B1          // 1
	C1 = "c"    // c
	D1          // c
	E1 = iota   // 4
	F1          // 5
)

func main() {

	println(a, b)
	println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	println(KB, MB, GB, TB)
	println(A, B, C, D)
	println(A1, B1, C1, D1, E1, F1)

	c := Black
	test(c)
	test(Red)
	test(Blue)

/*	x := 1
	test03(x) // cannot use x (type int) as type Color in argument to test03*/
	test(1) // 常量会被编译器自动转换

	a, b, c, d := 071, 0x1F, 1e9, math.MinInt16
	println(a, b, c, d)
}

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color)  {
}




