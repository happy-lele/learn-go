package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	t := make([]int, len(s), (cap(s) + 1) * 2)
	/*for i := range s {
		t[i] = s[i]
	}
	s = t*/
	//func copy(dst, src []T) int 把src复制到dst
	fmt.Println(copy(t, s))
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t,len(t), cap(t))

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)

	slice := make([]string, 0)
	fmt.Printf("切片的长度为%d 容量为%d\n", len(slice), cap(slice))
}

// AppendByte 在切片的末尾追加数据
func AppendByte(slice []byte, data ...byte) []byte{
	fmt.Printf("切片的长度为%d 容量为%d\n", len(slice), cap(slice))
	m := len(slice)
	n := m + len(data)
	// 需要的容量要大于现有的容量,n + 1防止n为0的情况
	if n > cap(slice) {
		newSlice := make([]byte, (n + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	fmt.Printf("切片的长度为%d 容量为%d\n", len(slice), cap(slice))
	slice = slice[0: n]
	fmt.Printf("切片的长度为%d 容量为%d\n", len(slice), cap(slice))
	copy(slice[m :n], data)
	return slice
}