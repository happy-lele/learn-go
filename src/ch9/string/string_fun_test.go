package string_fn

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	// 整形转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)

	// 字符串转整形
	if i, err:=strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
} 