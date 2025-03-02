package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}
}

func pow(x, n, lim float64) float64 {
	// if statement; 조건문{} :조건을 검사하기 전에 간단한 statement 실행
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
