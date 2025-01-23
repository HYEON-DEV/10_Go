package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// 사용자 데이터를 저장할 맵
var users = map[string]*User{}

// GET 요청을 처리하는 핸들러 함수
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// POST 요청을 처리하는 핸들러 함수
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	users[user.Nickname] = &user
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter()

	// GET 및 POST 요청을 처리하는 엔드포인트 설정
	router.HandleFunc("/api/get/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/api/post/users", CreateUserHandler).Methods("POST")

	log.Println("Starting server on :8081")
	http.ListenAndServe(":8081", router)
}
