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
		first10NumberSum       int
		evenNumbersSum         int
		oddNumbersSum          int
		n                      int
		numberBeforeLastNumber int
		lastNumber             int
	)

	valid = false

	if len(tckNumber) != 11 {
		err = errInvalidLength

		return
	}

	if tckNumber[0] == '0' {
		err = errInvalidIdentity
		return
	}

	numberBeforeLastNumber, err = strconv.Atoi(string(tckNumber[9]))

	if err != nil {
		return
	}

	lastNumber, err = strconv.Atoi(string(tckNumber[10]))

	if err != nil {
		return
	}

	for i := 0; i < len(tckNumber); i++ {

		n, err = strconv.Atoi(string(tckNumber[i]))

		if err != nil {
			return
		}

		if i != 10 {
			first10NumberSum += n
		}

		if i%2 == 0 && i <= 8 {
			evenNumbersSum += n
		}

		if i%2 != 0 && i < 8 {
			oddNumbersSum += n
		}
	}

	isFirst10NumberSumEqualLastNumber := (first10NumberSum % 10) == lastNumber

	if isFirst10NumberSumEqualLastNumber == false {
		err = errInvalidSum
		return
	}

	isRuleTrue := ((evenNumbersSum*7 + oddNumbersSum*9) % 10) == numberBeforeLastNumber
	isEvenRuleTrue := (evenNumbersSum*8)%10 == lastNumber

	rule := (evenNumbersSum*7-oddNumbersSum)%10 == numberBeforeLastNumber
	isLastNumberEven := lastNumber%2 == 0

	if (isRuleTrue && isEvenRuleTrue && rule && isLastNumberEven) == false {
		err = errInvalidIdentityValidation
		return
	}

	return true, nil
}
