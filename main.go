package main

import "fmt"

// custom types
type transformFn func(int) int

func main() {
	numbers := []int{2, 3, 4, 5, 6}
	doubled := transformNumbers(&numbers, double)

	resultTwo := transformNumbers(&numbers, getTransformerFn(&numbers))

	fmt.Println(doubled)
	fmt.Println(resultTwo)
}

// function that returns another functions
func getTransformerFn(numbers *[]int) func(int) int {

	// (dereference before using the numbers slice)
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
