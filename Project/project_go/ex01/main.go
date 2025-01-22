package main

import "net/http" // HTTP 서버와 클라이언트를 구현하기 위해 필요한 패키지

func main() {

	/*
	* 루트 URL ("/")에 대한 요청을 처리하는 핸들러 함수 등록
	* 이 핸들러 함수는 두 개의 매개변수를 받는다
	* rw는 응답 작성기(http.ResponseWriter)이고,
	* r은 요청(*http.Request)
	* 원하는 경로를 함수와 연결시킬 수 있다
	 */
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello"))
	})

	/*
	* TCP 포트 8080에서 HTTP 서버 시작
	* 두 번째 매개변수는 기본 멀티플렉서를 사용하기 위해 nil로 설정된다
	 */
	http.ListenAndServe(":8080", nil)
}
