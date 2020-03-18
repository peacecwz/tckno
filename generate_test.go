package tckno

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			name: "Generate TCK Number and Validate",
			data: Generate(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := Validate(tt.data)
			if !valid || err != nil {
				t.Errorf("Generate() = %v, want %v", valid, tt.data)
			}
		})
	}
}
