package main

import (
	"fmt"
	"test"
)

func main() {
	fmt.Printf("hello World: %s", "")
	test.PublicMethod()
	// test.privateMethod() は呼べない
}