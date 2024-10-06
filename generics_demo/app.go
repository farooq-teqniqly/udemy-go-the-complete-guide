package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add(3.4, 4.2))
	fmt.Println(add("hello", "world"))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
