package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	message := []byte("Hello, Gophers!")
	filePath := "c:/HYEON_GitHub/10_Go/WriteFile/hello.txt" // 절대 경로로 파일 경로 지정

	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		log.Fatalf("파일 쓰기 실패: %v", err)
	}

	// 파일이 성공적으로 작성되었는지 확인
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("파일이 존재하지 않습니다: %v", err)
	} else {
		log.Println("파일이 성공적으로 작성되었습니다.")
	}
}
