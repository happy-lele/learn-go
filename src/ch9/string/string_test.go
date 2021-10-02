package string_fn_test

import "testing"

func TestString(t *testing.T)  {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(s)

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))

	// 取出字符串里的unicode，go的一个内置机制，帮我们做这样一个转换
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}

	for _, c := range s {
		t.Logf("%[1]c %[1]c", c)
	}
}