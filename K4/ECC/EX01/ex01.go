package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

// ECC256 구조체 정의
// type ECC256 struct {
//     privateKey *ecdsa.PrivateKey
//     publicKey  ecdsa.PublicKey
// }

// 키 쌍 생성 메서드
// func (e *ECC256) GenerateKeys() error {
//     var err error
//     e.privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
//     if err != nil {
//         return err
//     }
//     e.publicKey = e.privateKey.PublicKey
//     return nil
// }

// // 서명 생성 메서드
// func (e *ECC256) SignMessage(message string) ([]byte, error) {
//     hash := sha256.Sum256([]byte(message))
//     r, s, err := ecdsa.Sign(rand.Reader, e.privateKey, hash[:])
//     if err != nil {
//         return nil, err
//     }
//     return append(r.Bytes(), s.Bytes()...), nil
// }

// // 서명 검증 메서드
// func (e *ECC256) VerifySignature(message string, signature []byte) bool {
//     hash := sha256.Sum256([]byte(message))
//     r := new(big.Int).SetBytes(signature[:len(signature)/2])
//     s := new(big.Int).SetBytes(signature[len(signature)/2:])
//     return ecdsa.Verify(&e.publicKey, hash[:], r, s)
// }

// func main() {
//     ecc := &ECC256{}

//     // 키 쌍 생성
//     err := ecc.GenerateKeys()
//     if err != nil {
//         log.Fatalf("키 생성 실패: %v", err)
//     }

//     message := "hello, world"

//     // 서명 생성
//     signature, err := ecc.SignMessage(message)
//     if err != nil {
//         log.Fatalf("서명 생성 실패: %v", err)
//     }
//     fmt.Printf("서명: %x\n", signature)

//     // 서명 검증
//     valid := ecc.VerifySignature(message, signature)
//     fmt.Printf("서명 검증 결과: %v\n", valid)
// }

// ECC256 구조체 정의
type ECC256 struct {
	Secp256r1 *ecdsa.PrivateKey
	P256      ecdsa.PublicKey
}

// 서명 생성 메서드
func sign(ecc *ECC256, message string) ([]byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	ecc.Secp256r1 = privateKey
	ecc.P256 = privateKey.PublicKey

	hash := sha256.Sum256([]byte(message))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// 서명 검증 메서드
func verify(ecc *ECC256, message string, signature []byte) bool {
	hash := sha256.Sum256([]byte(message))
	return ecdsa.VerifyASN1(&ecc.P256, hash[:], signature)
}

func main() {
	ecc := &ECC256{}

	message := "hello, world"

	// 서명 생성
	signature, err := sign(ecc, message)
	if err != nil {
		log.Fatalf("서명 생성 실패: %v", err)
	}
	fmt.Printf("서명: %x\n", signature)

	// 서명 검증
	valid := verify(ecc, message, signature)
	fmt.Printf("서명 검증 결과: %v\n", valid)
}
