package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,4,5}
	arr1[1] = 5
	arr4 := [2][2]int{{1,2}, {3,4}}

	t.Log(arr[1], arr[2])
	t.Log(arr1)
	t.Log(arr3)
	t.Log(arr4)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	arr3_sec := arr3[:3]  // 取前三个元素
	arr4_sec := arr3[3:]  // 下标从三开始的所有元素
	arr5_sec := arr3[:]  //  所有的元素
	t.Log(arr3_sec)
	t.Log(arr4_sec)
	t.Log(arr5_sec)
}

