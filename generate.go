package tckno

import (
	"strconv"
)

func Generate() string {

	var (
		digits     [11]int
		tcIdentity string
	)

	digits[0] = Random(1, 10)
	for i := 1; i < 9; i++ {
		digits[i] = Random(0, 10)
	}

	for i := 0; i < 9; i++ {
		digits[9] += digits[i]
	}

	evenTotal := digits[0] + digits[2] + digits[4] + digits[6] + digits[8]
	oddTotal := digits[1] + digits[3] + digits[5] + digits[7]

	digits[9] = (evenTotal*7 - oddTotal) % 10
	digits[10] = (evenTotal + oddTotal + digits[9]) % 10

	for i := 0; i < len(digits); i++ {
		tcIdentity += strconv.Itoa(digits[i])
	}

	return tcIdentity
}
