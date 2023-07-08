package bankCard

import (
	"unicode"
)

const (
	ten          = 10
	countNumbers = 16
)

func isValidCardNumber(cardNumber string) bool {
	numbers := make([]int, 0)
	for _, char := range cardNumber {
		if unicode.IsDigit(char) {
			numbers = append(numbers, int(char-'0'))
		}
	}
	if len(numbers) != countNumbers {
		return false
	}
	var sum int
	for ind, number := range numbers {
		n := number
		if ind%2 == 0 {
			n *= 2
		}
		if n > 9 {
			n -= 9
		}
		sum += n
	}
	return sum%ten == 0
}
