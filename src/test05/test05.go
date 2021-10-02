package main

func main() {
	s := "abcd"
	bs := []byte(s)

	bs[1] = 'B'
	println(string(bs))

	u := "电脑"
	us := []rune(u)

	us[1] = '话'
	println(string(us))
}
