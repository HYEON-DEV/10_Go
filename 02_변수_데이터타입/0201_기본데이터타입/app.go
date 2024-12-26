// 동일 폴더 내의 같은 패키지는 하나의 실행 파일로 묶인다.
// 이 때, main함수는 하나만 존재해야 하므로 폴더를 구분하여 main의 중복 선언 문제를 해결한다.
package main

import "fmt"

func main() {
	var unitName string = "Marine"
	var unitHP int = 40
	var unitAttack uint = 6
	var unitArmor uint = 0
	var attackSpeed float32 = 0.86
	var isFlying bool = false

	fmt.Println("유닛 정보:")
	fmt.Println("이름:", unitName)
	fmt.Println("체력:", unitHP)
	fmt.Println("공격력:", unitAttack)
	fmt.Println("방어력:", unitArmor)
	fmt.Println("공격 속도:", attackSpeed)
	fmt.Println("공중 유닛 여부:", isFlying)
}
