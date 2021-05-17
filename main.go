package main

import "fmt"

func main() {
	ans := Foo(1, 2)
	fmt.Println(ans)
}

func Foo(a, b int) int {
	return a + b
}
