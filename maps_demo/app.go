package main

import "fmt"

func main() {
	ages := map[string]int{
		"FM": 47,
		"MJ": 43,
		"NM": 17,
		"YM": 15,
		"LG": 19,
	}

	fmt.Println(ages)
	fmt.Println(ages["FM"])

	ages["Bubba"] = 38
	fmt.Println(ages)

	delete(ages, "Bubba")

	value, ok := ages["Bubba"]

	if !ok {
		fmt.Println("Bubba is not in the map")
	} else {
		fmt.Println(value)
	}

	fmt.Println(ages)
}
