package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // 切片的声明
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

//切⽚共享存储结构
func TestSliceShareMemory(t *testing.T){
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Juy", "Aug", "Sep", "Oct", "Nov", "dec"}

	Q2 := year[3:6]  // 取下标为 3~5 之间的所有元素
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) // [Jun Juy Aug] 3 7

	summer[0] = "Unknown"
	t.Log(Q2) // 无论被谁修改，都会输出修改后的结果
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}

	b := make([]int, 2, 4)

	t.Log(a, b)
/*	if a == b {
		t.Log("true")
	}*/
}