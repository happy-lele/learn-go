package try_test

import (
	"fmt"
	"testing"
)

var a int
func TestFirstTry(t *testing.T) {
	//var a int = 1
	//var b int = 2
	//var (
	//	b int = 1
	//	a     = 1
	//)
	//a := 1
	a = 1
	b := 1
	t.Log(a, "")
	for i:=0; i < 5; i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchange(t *testing.T) {
	var a int = 1
	var b int = 2
	a, b = b, a

	t.Log(a, b)
}