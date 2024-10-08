package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add(3.4, 4.2))
	fmt.Println(add("hello", "world"))

	intNumbers := []int{1, 2, 3, 4, 5}

	transform(&intNumbers, multiply(10))

	floatNumbers := []float64{1.3, 2.3, 3.3, 4.3, 5.3}

	transform(&floatNumbers, multiply(10.5))

	fmt.Println(intNumbers)
	fmt.Println(floatNumbers)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func multiply[T int | float64](factor T) func(T) T {
	return func(val T) T {
		return val * factor
	}
}

func transform[T int | float64](numbers *[]T, transformer func(T) T) {
	for i, val := range *numbers {
		transformed := transformer(val)
		(*numbers)[i] = transformed
	}
}
