package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
)

/* = = = = = = = JSON 형식으로 저장할 구조체 정의 = = = = = = = */

type PrivateKey struct {
	D     *big.Int // 개인 키의 비밀 정수 D
	X     *big.Int // 공개 키의 X 좌표
	Y     *big.Int // 공개 키의 Y 좌표
	Curve string   // 사용된 타원 곡선의 이름
}

type PublicKey struct {
	X     *big.Int
	Y     *big.Int
	Curve string
}

func main() {

	//  P-256 타원 곡선을 사용하여 새로운 ECDSA 개인 키 생성.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("ECDSA 개인키 생성 실패: %v", err) // %v: 값을 기본 형식으로 출력
	}

	/* = = = = = = = 키를 JSON 형식으로 변환 = = = = = = = */

	jsonPrivateKey := &PrivateKey{
		D:     privateKey.D,
		X:     privateKey.X,
		Y:     privateKey.Y,
		Curve: "P-256",
	}

	jsonPublicKey := &PublicKey{
		X:     privateKey.PublicKey.X,
		Y:     privateKey.PublicKey.Y,
		Curve: "P-256",
	}

	/* = = = = = = = 키를 JSON으로 직렬화 = = = = = = = */

	privateKeyData, err := json.Marshal(jsonPrivateKey)
	if err != nil {
		log.Fatalf("개인키 JSON 직렬화 실패: %v", err)
	}

	publicKeyData, err := json.Marshal(jsonPublicKey)
	if err != nil {
		log.Fatalf("공개키 JSON 직렬화 실패: %v", err)
	}

	/* = = = = = = = 직렬화된 JSON 데이터를 파일에 저장 = = = = = = = */

	err = os.WriteFile("c:/HYEON_GitHub/10_Go/K4/KeyFiles/PrivateKey.json", privateKeyData, 0644)
	if err != nil {
		log.Fatalf("개인키 파일 저장 실패: %v", err)
	}
	fmt.Println("개인키 파일 생성 - 성공")

	err = os.WriteFile("c:/HYEON_GitHub/10_Go/K4/KeyFiles/PublicKey.json", publicKeyData, 0644)
	if err != nil {
		log.Fatalf("공개키 파일 저장 실패: %v", err)
	}
	fmt.Println("공개키 파일 생성 - 성공")
}

/*
json.Marshal
Go의 encoding/json 패키지에서 제공하는 함수로,
Go 구조체를 JSON 형식의 바이트 슬라이스로 변환(직렬화)하는 데 사용
이 함수는 Go 구조체를 JSON 문자열로 변환하여 네트워크 전송이나 파일 저장 등에 사용할 수 있게 한한다.
*/

/*
직렬화(Serialization)
데이터 구조나 객체 상태를 저장하거나 전송할 수 있는 형식으로 변환하는 과정
직렬화된 데이터는 파일에 저장하거나 네트워크를 통해 전송할 수 있으며,
나중에 역직렬화(Deserialization)를 통해 원래의 데이터 구조나 객체 상태로 복원할 수 있다.
*/
