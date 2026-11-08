package main

import (
	"fmt"
)

func test(a int, b string, c int, d int) string {
	return fmt.Sprintf("%d %s %d %d", a, b, c, d)
}

func test3(a, b, c, d int) int {
	return a + b + c - d
}

func tren(a int, c int) (int, int) {
	sum := a + c
	multion := a * c

	return sum, multion
}
func tren4(a int, c int, d float64, b int) (int, int, float64, int) {

	sum := a + c
	multion := a * c
	ma := float64(int(d) % b)
	re := c / b

	return sum, multion, ma, re
}

func main() {
	fmt.Println("Hello")
	point := test(23, "wegdfy", 777, 5)
	fmt.Println(point)

	sum, multion := tren(23, 56)
	sum, multion, ma, ra := tren4(23, 56, 90.877, 676)
	fmt.Println("Sum:", sum)
	fmt.Println("Multiplication:", multion)
	fmt.Println("Modulus:", ma)
	fmt.Println("Division:", ra)

	fmt.Println(sum, multion)

	// Anonymous-Function

	func() {
		fmt.Println("Hello Anonymous-Function")
	}()

	// Variable Anonymous-Function
	new := func(a, b int) int {
		return a + b
	}

	fmt.Println(new(200, 100))
}
