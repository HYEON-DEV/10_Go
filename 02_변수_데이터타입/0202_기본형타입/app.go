package main

import "fmt"

func main() {
	var marineAttack int8 = 6      // 마린 공격력 (int8)
	var marineHP int16 = 40        // 마린 체력 (int16)
	var zerglingAttack int8 = 5    // 저글링 공격력 (int8)
	var zerglingHP int16 = 35      // 저글링 체력 (int16)
	var siegeTankAttack int32 = 70 // 시즈탱크 공격력 (int32)
	var siegeTankHP int64 = 150    // 시즈탱크 체력 (int64)

	fmt.Println("유닛 기본 속성: ")
	fmt.Println("마린 - 공격력: ", marineAttack, "\t체력: ", marineHP)
	fmt.Println("저글링 - 공격력:", zerglingAttack, "\t체력:", zerglingHP)
	fmt.Println("시즈탱크 - 공격력:", siegeTankAttack, "\t체력:", siegeTankHP)

	// 복소수 complex
	var marinePosition complex64 = 5 + 2i      // 마린의 위치
	var zerglingPosition complex64 = 3 + 4i    // 저글링의 위치
	var siegeTankPosition complex128 = 7 + 10i // 시즈탱크의 위치

	fmt.Println("\n유닛 위치:")
	fmt.Println("마린 위치:", marinePosition)
	fmt.Println("저글링 위치:", zerglingPosition)
	fmt.Println("시즈탱크 위치:", siegeTankPosition)

	// 유니코드 아이콘 (문자 데이터 타입)  rune
	var marineIcon rune = '⚔'    // 마린 아이콘 (검 모양)
	var zerglingIcon rune = '🐜'  // 저글링 아이콘 (개미 모양)
	var siegeTankIcon rune = '🚜' // 시즈탱크 아이콘 (탱크 상징)

	fmt.Println("\n유닛 아이콘:")
	fmt.Println("마린 아이콘:", string(marineIcon))
	fmt.Println("저글링 아이콘:", string(zerglingIcon))
	fmt.Println("시즈탱크 아이콘:", string(siegeTankIcon))

	// 포인터 크기 데이터 타입
	var mapIdentifier uintptr = 12345678 // 맵의 고유 식별자

	fmt.Println("\n맵 정보:")
	fmt.Println("맵 식별자:", mapIdentifier)
}
