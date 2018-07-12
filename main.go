package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//Verify("12345678901")
	fmt.Printf(Generate())
}

func Validate(strTcIdentity string) bool {
	if len(strTcIdentity) != 11 {
		fmt.Printf("Invalid length")
		return false
	}

	if strTcIdentity[0] == '0' {
		fmt.Printf("Invalid Identity")
		return false
	}

	tcIdentity, err := strconv.Atoi(strTcIdentity)
	if err != nil {
		log.Panic(err)
		return false
	}

	var (
		first10NumberSum int
		evenNumbersSum   int
		oddNumbersSum    int
	)

	for i := 0; i < len(strTcIdentity); i++ {

		n, err := strconv.Atoi(string(strTcIdentity[i]))
		if err != nil {
			log.Panic(err)
			return false
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

	numberBeforeLastNumber, err := strconv.Atoi(string(strTcIdentity[9]))
	if err != nil {
		log.Panic(err)
		return false
	}

	lastNumber, err := strconv.Atoi(string(strTcIdentity[10]))
	if err != nil {
		log.Panic(err)
		return false
	}

	isFirst10NumberSumEqualLastNumber := ((first10NumberSum % 10) == lastNumber)

	if isFirst10NumberSumEqualLastNumber == false {
		fmt.Printf("Invalid sum of Identity equal")
		return false
	}

	isRuleTrue := ((evenNumbersSum*7 + oddNumbersSum*9) % 10) == numberBeforeLastNumber
	isEvenRuleTrue := (evenNumbersSum*8)%10 == lastNumber

	rule := (evenNumbersSum*7-oddNumbersSum)%10 == numberBeforeLastNumber
	isLastNumberEven := lastNumber%2 == 0

	if (isRuleTrue && isEvenRuleTrue && rule && isLastNumberEven) == false {
		fmt.Printf("Invalid Identity Validation")
		return false
	}

	fmt.Printf(strconv.Itoa(tcIdentity) + " is valid")
	return true
}

func Verify() {

}

func Generate() string {
	rand.Seed(time.Now().Unix())
	var (
		digits     [11]int
		tcIdentity string
	)

	digits[0] = random(1, 10)
	for i := 1; i < 9; i++ {
		digits[i] = random(0, 10)
	}

	for i := 0; i < 9; i++ {
		digits[9] += digits[i]
	}
	//TODO: Refactoring
	digits[9] = ((digits[0]+digits[2]+digits[4]+digits[6]+digits[8])*7 -
		(digits[1] + digits[3] + digits[5] + digits[7])) % 10
	digits[10] = (digits[0] + digits[1] + digits[2] + digits[3] + digits[4] + digits[5] +
		digits[6] + digits[7] + digits[8] + digits[9]) % 10

	for i := 0; i < len(digits); i++ {
		tcIdentity += strconv.Itoa(digits[i])
	}

	return tcIdentity
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
