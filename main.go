package main

import "tc-identity-utils/tckno"

func main() {

	no := tckno.Generate()

	ok, err := tckno.Validate(no)

	if err != nil {
		panic(err)
	}

	if ok {
		println("validated!")
	} else {
		println("invalid tckno")
	}
}
