package main

import (
	"fmt"
	"testing"
)

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "不存在数据", false
}

func TestGetValue(t *testing.T) {
	intmap := map[int]string{
		1:"a",
		2:"b",
		3:"ccc",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}
