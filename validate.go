package tckno

import (
	"errors"
	"strconv"
)

var (
	errInvalidLength             = errors.New("Invalid length")
	errInvalidIdentity           = errors.New("Invalid TCK Number")
	errInvalidIdentityValidation = errors.New("Invalid Identity Validation")
	errInvalidSum                = errors.New("Invalid sum of Identity equal")
)

func Validate(tckNumber string) (valid bool, err error) {
	var (
		sumOfEvenDigits  int
		sumOfOddDigits   int
		digit            int
		penultimateDigit int
		lastDigit        int
	)

	if len(tckNumber) != 11 {
		err = errInvalidLength
		return
	}

	if tckNumber[0] == '0' {
		err = errInvalidIdentity
		return
	}

	penultimateDigit, err = strconv.Atoi(string(tckNumber[9]))
	if err != nil {
		return
	}

	lastDigit, err = strconv.Atoi(string(tckNumber[10]))
	if err != nil {
		return
	}

	for i, tckDigit := range tckNumber {
		digit, err = strconv.Atoi(string(tckDigit))
		if err != nil {
			return
		}

		if i%2 == 0 && i <= 8 {
			sumOfEvenDigits += digit
		}

		if i%2 != 0 && i < 8 {
			sumOfOddDigits += digit
		}
	}

	isPenultimateDigitValid := (sumOfEvenDigits*7-sumOfOddDigits)%10 == penultimateDigit
	if !isPenultimateDigitValid {
		err = errInvalidIdentityValidation
		return
	}

	isLastDigitValid := (sumOfEvenDigits+sumOfOddDigits+penultimateDigit)%10 == lastDigit
	if !isLastDigitValid {
		err = errInvalidSum
		return
	}

	return true, nil
}
