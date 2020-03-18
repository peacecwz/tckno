package tckno

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		tckNumber string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
		wantErr   bool
	}{
		{
			name: "Test valid TCK number",
			args: args{
				tckNumber: "49786915818",
			},
			wantErr:   false,
			wantValid: true,
		},
		{
			name: "Test has invalid character in TCK Number",
			args: args{
				tckNumber: "4978691581f",
			},
			wantErr:   true,
			wantValid: false,
		},
		{
			name: "Test has invalid length in TCK Number",
			args: args{
				tckNumber: "4978691581",
			},
			wantErr:   true,
			wantValid: false,
		},
		{
			name: "Test invalid TCK Number",
			args: args{
				tckNumber: "09786915812",
			},
			wantErr:   true,
			wantValid: false,
		},
		{
			name: "Test invalid TCK Number",
			args: args{
				tckNumber: "45786915812",
			},
			wantErr:   true,
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, err := Validate(tt.args.tckNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValid != tt.wantValid {
				t.Errorf("Validate() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
