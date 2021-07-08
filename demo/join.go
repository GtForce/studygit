package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"hello", "world", "!"}
	res := strings.Join(str1, " ")
	fmt.Printf("res:%s\n", res)

	res1 := bytes.Join([][]byte{[]byte("hello"), []byte("world"), []byte("!")}, []byte(" "))
	fmt.Printf("res1:%s\n", res1)
}
