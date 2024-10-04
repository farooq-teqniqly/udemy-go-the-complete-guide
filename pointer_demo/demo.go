package main

import (
	"errors"
	"fmt"
)

func main() {
	age := 47
	fmt.Println("Age before method call:", age)
	err := subtractLegalAge(&age)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Age after method call:", age)
}

// subtractLegalAge subtracts 18 from the given age, modifying the input pointer directly.
func subtractLegalAge(age *int) error {
	if age == nil {
		return errors.New("age is nil")
	}

	*age = *age - 18
	return nil
}
