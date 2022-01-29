package main

import (
	"testing"
)

func Test_validateBinary(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "match",
			args:    args{value: "0101"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not match",
			args:    args{value: "01234"},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateBinary(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertDecimal(t *testing.T) {
	type args struct {
		decimal int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{decimal: 0},
			want: "0",
		},
		{
			name: "9",
			args: args{decimal: 9},
			want: "1001",
		},
		{
			name: "19",
			args: args{decimal: 19},
			want: "10011",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDecimal(tt.args.decimal); got != tt.want {
				t.Errorf("convertDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertBinary(t *testing.T) {
	type args struct {
		binary string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "0",
			args:    args{binary: "0"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "9",
			args:    args{binary: "1001"},
			want:    9,
			wantErr: false,
		},
		{
			name:    "19",
			args:    args{binary: "10011"},
			want:    19,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertBinary(tt.args.binary)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
