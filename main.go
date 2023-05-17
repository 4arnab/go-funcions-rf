package main

import "fmt"

// custom types
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	doubled := transformNumbers(&numbers, double)

	fmt.Println(doubled)
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
