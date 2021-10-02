package main

import "unsafe"

func main() {
	s := "abc"
	println(&s)

	s, y := "hello", 20
	println(&s, y)

	{
		s, z := 1000, 30
		println(&s, z)
	}

	const a, b = 1, 2
	const (
		c = false
		d // false

		e = "abc"
		f = len(e)
		g = unsafe.Sizeof(f)

	)

	println(a, b, c, d, e, f, g)



}
