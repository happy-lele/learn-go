package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Println(s, len(s), cap(s))

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b, len(b), cap(b))
}
