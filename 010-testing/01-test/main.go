package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 7 = ", mySum(4, 7))
	fmt.Println("5 + 9 = ", mySum(5, 9))


	fmt.Println("2 * 3 = ", multiply(2, 3))
	fmt.Println("4 * 5 = ", multiply(4, 5))
	fmt.Println("3 * 6 = ", multiply(3, 6))


}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func multiply(xi ...int) int {
	total := 1
	for _, v := range xi {
		total *= v
	}
	return total
}