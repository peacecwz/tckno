package tckno

import "testing"

func TestRandom(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Generate random number",
			args: args{
				min: 1,
				max: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			random := Random(tt.args.min, tt.args.max)
			if random <= tt.args.min && random > tt.args.max {
				t.Errorf("Random() = %v, min %v max %v", random, tt.args.min, tt.args.max)
			}
		})
	}
}
