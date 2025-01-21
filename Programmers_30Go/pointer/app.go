package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // i를 가리키는 포인터
	fmt.Println(*p) // 포인터를 통해 i값을 읽는다

	*p = 21 // 포인터를 통해 i값을 설정한다
	fmt.Println(i)

	p = &j   // j를 가리킨다
	*p /= 37 // 포인터를 통해 j를 나눈다
	fmt.Println(j)
}
