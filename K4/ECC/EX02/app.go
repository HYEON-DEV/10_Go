package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type ECC256 struct {
	Secp256r1 *ecdsa.PrivateKey
	P256      ecdsa.PublicKey
}

func sign(ecc *ECC256, signature string) ([]byte, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	ecc.Secp256r1 = privateKey
	ecc.P256 = privateKey.PublicKey
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte(signature))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}

	return sig, hash[:]
}

func verify(ecc *ECC256, signature string) bool {
	hash, sig := sign(ecc, signature)
	valid := ecdsa.VerifyASN1(&ecc.Secp256r1.PublicKey, hash[:], sig)
	return valid
}

func main() {
	ecc := &ECC256{}
	signature := "hyeon"
	sig, _ := sign(ecc, signature)
	fmt.Printf("Signature: %x\n", sig)
	fmt.Println("verify: ", verify(ecc, signature))
}
