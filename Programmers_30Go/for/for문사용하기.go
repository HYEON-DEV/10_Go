package main

import "fmt"

func main() {
	// ex 1)
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// ex 2)
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
