package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
)

type PrivateKey struct {
	D     *big.Int
	X     *big.Int
	Y     *big.Int
	Curve string
}

func main() {
	// 개인 키 파일 경로
	privateKeyFilePath := "c:/HYEON_GitHub/10_Go/K4/KeyFiles/PrivateKey.json"

	// 개인 키 파일 읽기기
	privateKeyData, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatalf("개인 키 파일 읽기 실패: %v", err)
	}

	// JSON 형식의 개인 키 데이터를 PrivateKey 구조체로 변환
	var jsonPrivateKey PrivateKey
	err = json.Unmarshal(privateKeyData, &jsonPrivateKey)
	if err != nil {
		log.Fatalf("개인 키 JSON 파싱 실패: %v", err)
	}

	// 타원 곡선 P-256 설정
	curve := elliptic.P256()
	privateKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     jsonPrivateKey.X,
			Y:     jsonPrivateKey.Y,
		},
		D: jsonPrivateKey.D,
	}

	// 서명할 메시지 설정
	message := "hyeon"
	hash := sha256.Sum256([]byte(message))

	// ECDSA 서명 생성
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Fatalf("서명 생성 실패: %v", err)
	}

	// 서명 결과를 바이트 배열로 결합
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("서명: %x\n", signature)

	// err = os.WriteFile("Signature.txt", signature, 0644)
	// if err != nil {
	// 	log.Fatalf("서명 파일 저장 실패: %v", err)
	// }

	// fmt.Println("Signature saved to Signature.txt")
}
