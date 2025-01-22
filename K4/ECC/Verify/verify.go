package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
)

type PublicKey struct {
	X     *big.Int
	Y     *big.Int
	Curve string
}

func main() {
	// 공개 키 파일 경로
	publicKeyFilePath := "c:/HYEON_GitHub/10_Go/K4/KeyFiles/PublicKey.json"

	// 서명 파일 경로
	signatureFilePath := "c:/HYEON_GitHub/10_Go/K4/ECC/GenerateSign/Signature.txt"

	// 공개 키 파일 읽기
	publicKeyData, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		log.Fatalf("공개 키 파일 읽기 실패: %v", err)
	}

	// JSON 형식의 공개 키 데이터를 PublicKey 구조체로 변환
	var jsonPublicKey PublicKey
	err = json.Unmarshal(publicKeyData, &jsonPublicKey)
	if err != nil {
		log.Fatalf("공개 키 JSON 파싱 실패: %v", err)
	}

	// 타원 곡선 P-256 설정
	curve := elliptic.P256()
	publicKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     jsonPublicKey.X,
		Y:     jsonPublicKey.Y,
	}

	// 서명할 메시지 설정
	var message string
	fmt.Print("서명 메시지를 입력하세요: ")
	_, err = fmt.Scanln(&message)
	if err != nil {
		log.Fatalf("서명 메시지 입력 실패: %v", err)
	}
	hash := sha256.Sum256([]byte(message))

	// 서명 파일 읽기
	signatureData, err := os.ReadFile(signatureFilePath)
	if err != nil {
		log.Fatalf("서명 파일 읽기 실패: %v", err)
	}

	// 서명 데이터를 바이트 배열로 변환
	signature := signatureData

	// 서명 검증
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])
	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Println("서명 검증 결과:", valid)
}
