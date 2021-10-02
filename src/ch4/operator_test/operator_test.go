package operator_test

import "testing"

const(
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 5}

	t.Log(a == b)
	t.Log(a == d)
}

func TestBitClear(t *testing.T)  {
	a := 7 // 0111
	a = a &^ Readable  // 0111 &^ 0001 == 0110
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}