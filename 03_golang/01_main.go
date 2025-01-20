// 패키지가 main이면 전체 패키지 이고  main 패키지 안에 main 함수가 있어야 한다
package main

// 라이브러리 import
// PC 에 설치된 기본 라이브러리
import (
	"fmt"
	"math"
)

func main() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

	fmt.Print("The vector (", a, b, ") has length", h, ".\n")
	fmt.Println("The vector (", a, b, ") has length", h, ".")
	fmt.Printf("The vector (%g %g) has length %g. \n", a, b, h)

}
