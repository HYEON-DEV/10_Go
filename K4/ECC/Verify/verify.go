package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

type JSONPublicKey struct {
	X     *big.Int
	Y     *big.Int
	Curve string
}

func main() {
	publicKeyFilePath := "PublicKey.json"
	publicKeyData, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		fmt.Println("공개 키 파일 읽기 실패:", err)
		return
	}

	var jsonPublicKey JSONPublicKey
	err = json.Unmarshal(publicKeyData, &jsonPublicKey)
	if err != nil {
		fmt.Println("공개 키 JSON 파싱 실패:", err)
		return
	}

	curve := elliptic.P256()
	publicKey := &ecdsa.PublicKey{
		Curve: curve,
		X:     jsonPublicKey.X,
		Y:     jsonPublicKey.Y,
	}

	signature, err := os.ReadFile("Signature.txt")
	if err != nil {
		fmt.Println("서명 파일 읽기 실패:", err)
		return
	}

	message := "hyeon"
	hash := sha256.Sum256([]byte(message))

	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Println("서명 검증 결과:", valid)
}
