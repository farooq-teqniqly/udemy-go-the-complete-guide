package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add(3.4, 4.2))
	fmt.Println(add("hello", "world"))

	intNumbers := []int{1, 2, 3, 4, 5}

	transform(&intNumbers, multiply)

	floatNumbers := []float64{1.3, 2.3, 3.3, 4.3, 5.3}

	transform(&floatNumbers, multiply)

	fmt.Println(intNumbers)
	fmt.Println(floatNumbers)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func multiply[T int | float64](a T) T {
	return a * 2
}

func transform[T int | float64](numbers *[]T, transformer func(T) T) {
	for i, val := range *numbers {
		transformed := transformer(val)
		(*numbers)[i] = transformed
	}
}
