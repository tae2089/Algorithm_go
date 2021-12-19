package checksum

import (
	"strconv"
)

// 이 알고리즘은 luhn 알고리즘입니다.
// 관련 설명은 여기를 참조해주세요:)

func LuhnAlgorithm(s []rune) bool {

	number := 0
	total := 0
	for i, value := range s {
		number, _ = strconv.Atoi(string(value))
		if i%2 == 1 {
			total += number
		} else {
			number *= 2
			if number > 9 {
				number -= 9
			}
			total += number
		}
	}
	return total%10 == 0
}
