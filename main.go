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

	// anonymous functions
	transformed := transformNumberTwo(&numbers, func(number int) int {
		return number * 2
	})

	double := createTransformer(10)(10)

	fmt.Println(transformed, double)
}

// closures <it returns a closure function or anonymous function>
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumberTwo(number *[]int, fn func(int) int) []int {
	dNumbers := make([]int, 30)
	for _, value := range *number {
		dNumbers = append(dNumbers, fn(value))
	}

	return dNumbers
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
