package main

import (
	"crypto/sha256" // SHA-256 해시 알고리즘 제공하는 패키지
	"fmt"
)

// "hello world\n" 문자열의 SHA-256 해시 값이 16진수로 출력
func main() {
	// 문자열 "hello world\n"을 바이트 슬라이스로 변환
	// ( SHA-256 해시 함수는 바이트 슬라이스를 입력으로 받는다 )
	// sha256.Sum256: 주어진 데이터의 SHA-256 해시 값을 계산하는 함수
	sum := sha256.Sum256([]byte("hello world\n")) // 32byte배열

	// 바이트 슬라이스나 배열을 16진수로 출력하는 형식 지정자
	fmt.Printf("%x", sum)
}
