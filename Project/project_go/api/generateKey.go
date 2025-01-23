package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gorilla/mux"
)

// PrivateKey 구조체 정의
type PrivateKey struct {
	D     *big.Int `json:"d"`
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}

// PublicKey 구조체 정의
type PublicKey struct {
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
	Curve string   `json:"curve"`
}

var privateKey *ecdsa.PrivateKey
var publicKey ecdsa.PublicKey

// 키 생성 요청을 처리하는 핸들러 함수
func GenerateKey(w http.ResponseWriter, r *http.Request) {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Printf("ECDSA 개인키 생성 실패: %v", err)
		http.Error(w, "키 생성 실패", http.StatusInternalServerError)
		return
	}
	log.Println("ECDSA 키 생성")
	// publicKey = privateKey.PublicKey

	// 키를 JSON 형식으로 변환
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

	// 키를 JSON으로 직렬화

	privateKeyData, err := json.Marshal(jsonPrivateKey)
	if err != nil {
		log.Printf("개인키 JSON 직렬화 실패: %v", err)
		http.Error(w, "개인키 JSON 직렬화 실패", http.StatusInternalServerError)
		return
	}

	publicKeyData, err := json.Marshal(jsonPublicKey)
	if err != nil {
		log.Printf("공개키 JSON 직렬화 실패: %v", err)
		http.Error(w, "공개키 JSON 직렬화 실패", http.StatusInternalServerError)
		return
	}

	// 직렬화된 JSON 데이터를 파일에 저장

	// 사용자의 다운로드 폴더 경로 가져오기
	var downloadPath string
	if runtime.GOOS == "windows" {
		downloadPath = filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
	} else {
		downloadPath = filepath.Join(os.Getenv("HOME"), "Downloads")
	}
	privateKeyFile := filepath.Join(downloadPath, "PrivateKey.json")
	publicKeyFile := filepath.Join(downloadPath, "PublicKey.json")

	err = os.WriteFile(privateKeyFile, privateKeyData, 0644)
	if err != nil {
		log.Printf("개인키 파일 저장 실패: %v", err)
		http.Error(w, "개인키 파일 저장 실패", http.StatusInternalServerError)
		return
	}
	log.Println("개인키 파일 생성 - 성공")

	err = os.WriteFile(publicKeyFile, publicKeyData, 0644)
	if err != nil {
		log.Printf("공개키 파일 저장 실패: %v", err)
		http.Error(w, "공개키 파일 저장 실패", http.StatusInternalServerError)
		return
	}
	log.Println("공개키 파일 생성 - 성공")

	// 응답으로 성공 메시지 반환
	response := map[string]string{
		"message":        "키 생성이 완료되었습니다.",
		"privateKeyFile": privateKeyFile,
		"publicKeyFile":  publicKeyFile,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	// POST 요청을 처리하는 엔드포인트 설정
	router.HandleFunc("/api/generate_key", GenerateKey).Methods("POST")

	log.Println("Starting server on :8081")
	http.ListenAndServe(":8081", router)
}
