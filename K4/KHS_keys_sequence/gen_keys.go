package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	// GenerateKey 함수는 주어진 타원 곡선과 난수 생성기를 사용하여 비밀키와 공개키를 생성한다.
	// P256은 타원 곡선 암호화(Elliptic Curve Cryptography)의 한 종류로, 256비트 크기의 타원 곡선을 사용한다.
	// rand.Reader는 암호학적으로 안전한 난수를 생성하는데 사용되는 인터페이스로, 해당 인터페이스를 구현한 객체(=난수)를 인증키 생성에 사용한다.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // GenerateKey() 함수는 두 개의 값을 반환하는데, 첫 번째 값은 비밀키, 두 번째 값은 오류 객체이다.
	// privateKey는 구조체임

	if err != nil {
		panic(err) // 복구할 수 없는 에러가 발생했을 때, 프로그램의 중단과 고루틴을 패닉 상태로 만드는 함수.
	} // 프로그램은 즉시 실행을 중단하고, 현재 함수에서부터 호출 스택을 따라 올라가면서 모든 defer 함수들을 실행 수 프로그램 종료

	publicKey := &privateKey.PublicKey // 공개키를 추출한다.

	msg := "run doggo run"
	// Sum256() 함수는 주어진 바이트 슬라이스의 SHA-256 해시값을 계산하여 반환한다.
	// 이때, msg 문자열을 바이트 슬라이스로 변환하여 해시값을 계산한다.
	// 해쉬값을 만들때 쓰일 바이트 슬라이스는, r, u, n, , d, o, g, g, o, , r, u, n
	// 이걸 이용한 해쉬값을 계산.
	// 그 값을 hash에 저장.
	hash := sha256.Sum256([]byte(msg))

	sigR, sigS, err := ecdsa.Sign(rand.Reader, privateKey, hash[:]) // SignASN1() 함수 대신 Sign 함수로
	if err != nil {
		panic(err)
	}
	fmt.Printf("sigR = %x \nsigS= %x\npublicKey= %x", sigR, sigS, publicKey)

	sigRHex := fmt.Sprintf("%x", sigR)
	sigSHex := fmt.Sprintf("%x", sigS)

	// SignASN1에서 Sign으로 변경
	// pubKeyHex := fmt.Sprintf("%x", publicKey)

	// 공개 키를 16진수 문자열로 변환
	//pubKeyHex := fmt.Sprintf("%x", elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y))

	valid_1 := ecdsa.Verify(publicKey, hash[:], sigR, sigS)
	fmt.Println("Signature verified:", valid_1)

	// message_txt := []byte("&privateKey.PublicKey")
	//sigHex := fmt.Sprintf("%x", sig)

	// 원래 이 아래에 데이터를 텍스트 파일로 작성할 때,
	// JSON Marshal을 사용하여 데이터를 JSON 형식으로 변환한 뒤, 파일에 쓰는 것이 좋다.
	// 일단 지금은 단일 데이터 저장하는 것에 의의를 두고 간단하게 작성한다.

	err_txt := os.WriteFile("Private_key.txt", []byte(""), 0644) //0644는 8진수로 사용자 권한에 대한 의미.
	//err_txt + os.WriteFile("hello.txt", []byte(pubKeyHex), 0644)
	if err_txt != nil {
		log.Fatal(err_txt)
	}

	file_sig, err := os.OpenFile("Private_key.txt", os.O_APPEND|os.O_WRONLY, 0644) // os.O_APPEND|os.O_WRONLY는,
	if err != nil {                                                                // os.O_APPEND = 항상 열 때 파일 끝에 추가
		panic(err) // os.O_WRONLY = 파일을 쓰기 전용으로 열기
	}
	defer file_sig.Close() // 이 작업을 통해서 파일은 메인 함수가 종료되면 닫게 예약한다.

	_, err = file_sig.Write([]byte(sigRHex)) // 첫번째 반환값은 무시한다는 의미의 "_"
	if err != nil {
		panic(err)
	}

	_, err = file_sig.Write([]byte("\n"))
	if err != nil {
		panic(err)
	}

	_, err = file_sig.Write([]byte(sigSHex)) // 첫번째 반환값은 무시한다는 의미의 "_"
	if err != nil {
		panic(err)
	}

	// _, err = file.Write([]byte(pubKeyHex))
	// if err != nil {
	// 	panic(err)
	// }

	err_txt_pub := os.WriteFile("Public_key.txt", []byte(""), 0644) //0644는 8진수로 사용자 권한에 대한 의미.
	//err_txt + os.WriteFile("hello.txt", []byte(pubKeyHex), 0644)
	if err_txt_pub != nil {
		log.Fatal(err_txt_pub)
	}

	file_pub, err_pub := os.OpenFile("Public_key.txt", os.O_APPEND|os.O_WRONLY, 0644) // os.O_APPEND|os.O_WRONLY는,
	if err_pub != nil {                                                               // os.O_APPEND = 항상 열 때 파일 끝에 추가
		panic(err_pub) // os.O_WRONLY = 파일을 쓰기 전용으로 열기
	}
	defer file_pub.Close() // 이 작업을 통해서 파일은 메인 함수가 종료되면 닫게 예약한다.

	// elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y) 를 []byte(pubKeyHex) 대신 쓰면?
	_, err_pub = file_pub.Write(elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y)) // 첫번째 반환값은 무시한다는 의미의 "_"
	if err_pub != nil {
		panic(err_pub)
	}

}
