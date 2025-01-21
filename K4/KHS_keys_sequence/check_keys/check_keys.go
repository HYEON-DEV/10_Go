package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	// 파일에서 데이터 읽기
	data_priv, err_priv := os.ReadFile("Private_key.txt")
	if err_priv != nil {
		log.Fatal(err_priv)
	}

	// 파일 내용 파싱
	lines_priv := strings.Split(string(data_priv), "\n")
	if len(lines_priv) < 2 {
		log.Fatal("Invalid file format")
	}
	sigRHexStr := lines_priv[0]
	sigSHexStr := lines_priv[1]

	// 결과 출력
	fmt.Println("Parsed sigRHexStr:", sigRHexStr)
	fmt.Println("Parsed sigSHexByte:", sigSHexStr)

	data_pub, err_pub := os.ReadFile("Public_key.txt")
	if err_pub != nil {
		log.Fatal(err_pub)
	}

	// Public_key 파일 내용 파싱 --> 현재 일단 바이트 상태 그대로 저장받음
	// pubKeyHexStr := string(data_pub)
	pubKeyHex := data_pub // 이 코드는 바이트 상태 그대로 저장받았을 때 사용
	//
	// 결과 출력
	//fmt.Println("Parsed pubKeyHexStr:", pubKeyHexStr)
	//
	// sigRHex, sigSHex, pubKeyHex 값을 이용해서 검증해보기
	//pubKeyHex, err_pkh := hex.DecodeString(pubKeyHexStr)
	//if err_pkh != nil {
	//	log.Fatal(err_pkh)
	//}

	// 바이트 슬라이스를 공개 키로 변환
	x, y := elliptic.Unmarshal(elliptic.P256(), pubKeyHex)
	if x == nil || y == nil {
		log.Fatal("Invalid public key")
	}
	publicKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	// 16진수 문자열을 big.Int로 변환
	r := new(big.Int)
	r.SetString(sigRHexStr, 16)
	s := new(big.Int)
	s.SetString(sigSHexStr, 16)

	// 메시지 해시 계산
	msg := "run doggo run"
	hash := sha256.Sum256([]byte(msg))

	// 서명 검증
	valid := ecdsa.Verify(&publicKey, hash[:], r, s)
	fmt.Println("Signature verified:", valid)

}
