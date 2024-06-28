package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "test & pol"
	splittendText := strings.Split(text, " ")
	fmt.Println(len(splittendText))
	idx := strings.IndexByte(text, "&"[0])
	fmt.Println(idx, strings.TrimSpace(text[:idx]), strings.TrimSpace(text[idx+1:]))
}
