package tckno

import "testing"

func TestTCKNOValidation(t *testing.T) {

	validations := []struct {
		input string
	}{
		{"12345678"},
		{"00000000000"},
		{"12345678a00"},
		{"1234567890f"},
		{"12345678900"},
	}

	for _, tt := range validations {

		ok, err := Validate(tt.input)

		t.Log(err)

		if err == nil {
			t.Error("Validation success!")
		}

		if ok {
			t.Error("Invalid expected value:", tt.input)
		}

	}
}

func TestGenerate(t *testing.T) {
	vtckno := Generate()

	t.Log("tckno:", vtckno)

	ok, err := Validate(vtckno)

	if err != nil {
		t.Error("validation error:", err)
	}

	if !ok {
		t.Error("Cannot validate generated code")
	}
}
