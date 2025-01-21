package main

import (
	"crypto/ecdsa"    // ECDSA 알고리즘을 제공하는 패키지
	"crypto/elliptic" // 타원 곡선 연산을 제공하는 패키지
	"crypto/rand"     // 암호학적으로 안전한 난수를 생성하는 패키지
	"crypto/sha256"   // SHA-256 해시 알고리즘을 제공하는 패키지
	"fmt"
)

/** ECDSA(Elliptic Curve Digital Signature Algorithm)를 사용하여 메시지를 서명하고 서명을 검증 **/

// "hello, world"의 SHA-256 해시 값 계산
// ECDSA 개인 키를 사용하여 서명한 후,
// 서명의 유효성을 검증한다

func main() {
	// ecdsa.GenerateKey: 주어진 타원 곡선과 난수 생성기를 사용하여 새로운 ECDSA 개인 키를 생성하는 함수
	// elliptic.P256(): P-256 타원 곡선을 사용
	// rand.Reader: 암호학적으로 안전한 난수 생성기를 사용
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// 오류가 발생하면 프로그램을 중단하고 오류 메시지 출력
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	// sha256.Sum256: 주어진 데이터의 SHA-256 해시 값을 계산하는 함수
	hash := sha256.Sum256([]byte(msg))

	// ecdsa.SignASN1: 주어진 해시 값에 대해 ECDSA 서명을 생성하는 함수입니다.
	// ASN.1 인코딩 사용
	// hash[:]: 해시 값을 바이트 슬라이스로 전달
	// hash[:] : 배열이나 슬라이스의 전체 내용을 슬라이스로 변환
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	// ecdsa.VerifyASN1: 주어진 서명과 해시 값을 사용하여 서명의 유효성을 검증하는 함수.
	// ASN.1 인코딩 사용
	// &privateKey.PublicKey: 검증에 사용할 ECDSA 공개 키
	// 서명이 유효하면 true, 그렇지 않으면 false
	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)

}
