package tckno

import (
	"strconv"
)

func Generate() string {
	var (
		digits          [11]int
		tcIdentity      string
		sumOfEvenDigits int
		sumOfOddDigits  int
	)

	digits[0] = Random(1, 10)
	for i := 1; i < 9; i++ {
		digits[i] = Random(0, 10)
	}

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			sumOfEvenDigits += digits[i]
		} else {
			sumOfOddDigits += digits[i]
		}
	}

	digits[9] = (sumOfEvenDigits*7 - sumOfOddDigits) % 10
	digits[10] = (sumOfEvenDigits + sumOfOddDigits + digits[9]) % 10

	for _, digit := range digits {
		tcIdentity += strconv.Itoa(digit)
	}

	return tcIdentity
}
