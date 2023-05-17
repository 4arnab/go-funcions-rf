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

	// double := createTransformer(10)(10)
	fmt.Println(transformed)
	fmt.Println(factorial(5))
	fmt.Println(sumUp(1, 3, 4, 5, 5))
	fmt.Println(sumUp(numbers...)) // splitting up array or slice into variadic values
}

// defer keyword

// variadic function .<.. spread operator in js>
func sumUp(numbers ...int) (sum int) {
	// sum := 0
	for _, value := range numbers {
		sum += value
	}

	return
}

// factorial of 5 < 5 * 4 * 3 * 2 * 1>
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
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
