package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
* 문자열 키와 User 구조체 포인터 값을 가지는 맵을 선언하고 초기화
* 예를 들어, 키는 사용자 ID 또는 이메일 주소가 될 수 있다.
 */
var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"` // 문자열 타입이며, JSON으로 인코딩될 때 "nickname"이라는 이름으로 변환된다
	Email    string `json:"email"`    // 문자열 타입이며, JSON으로 인코딩될 때 "email"이라는 이름으로 변환된다
}

func main() {
	log.Println("Start server on :8081")

	http.HandleFunc("/users", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method { // 요청의 HTTP 메서드에 따라 다른 동작 수행
		case http.MethodGet:
			json.NewEncoder(wr).Encode(users) // GET 요청인 경우, users 맵을 JSON으로 인코딩하여 응답 본문에 쓴다다
		case http.MethodPost: // POST 요청인 경우, 요청 본문에서 JSON 데이터를 디코딩하여 user 변수에 저장
			var user User // users 맵에 새로운 사용자를 추가
			json.NewDecoder(r.Body).Decode(&user)
			users[user.Email] = &user
			json.NewEncoder(wr).Encode(user) // 추가된 사용자 정보를 JSON으로 인코딩하여 응답 본문에 쓴다
		}
	})
	http.ListenAndServe(":8081", nil)
}
