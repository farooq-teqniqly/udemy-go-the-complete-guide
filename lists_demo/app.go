package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	readings := createRandomReadings[float64](10)
	printReadings(readings)

	readings2 := createRandomReadings[int](10)
	printReadings(readings2)

	grades := []float64{85.0, 90.0, 96.0, 80.0}
	sort.Float64s(grades)
	fmt.Println(median(grades))

	grades = []float64{90.0, 80.0, 85.0}
	sort.Float64s(grades)
	fmt.Println(median(grades))

	grades2 := []int{85, 90, 96, 80}
	sort.Ints(grades2)
	fmt.Println(median(grades2))

	grades2 = []int{90, 80, 85}
	sort.Ints(grades2)
	fmt.Println(median(grades2))
}

func printReadings[T int | float64](readings []T) {
	readingsLength := len(readings)

	fmt.Print("[")

	for i := 0; i < readingsLength; i++ {
		switch any(readings[i]).(type) {
		case int:
			fmt.Print(readings[i])
		case float64:
			fmt.Printf("%.2f", readings[i])
		}

		if i < readingsLength-1 {
			fmt.Print(", ")
		}
	}

	fmt.Println("]")
}

func createRandomReadings[T int | float64](count int) []T {
	readings := make([]T, count)

	for i := 0; i < count; i++ {
		var value interface{}

		switch any(readings[i]).(type) {
		case int:
			value = rand.Intn(100)
		case float64:
			value = rand.Float64() * 100
		}

		readings[i] = value.(T)
	}

	return readings
}

func median[T int | float64](arr []T) T {
	readingsLength := len(arr)
	mid := readingsLength / 2

	if readingsLength%2 != 0 {
		return arr[mid]
	}

	return (arr[mid-1] + arr[mid]) / 2
}
