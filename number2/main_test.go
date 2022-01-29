package main

import (
	"reflect"
	"testing"
)

func Test_findPalindrome(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{text: "aku suka makan nasi"},
			want: "aku suka",
		},
		{
			name: "test 2",
			args: args{text: "di rumah saya ada kasur rusak"},
			want: "kasur rusak",
		},
		{
			name: "test 3",
			args: args{text: "abcde edcbza"},
			want: "bcde edcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPalindrome(tt.args.text); got != tt.want {
				t.Errorf("findPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSpaceIndex(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "aku suka makan nasi",
			args: args{text: "aku suka makan nasi"},
			want: []int{3, 8, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSpaceIndex(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSpaceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addSpace(t *testing.T) {
	source := "aku suka makan nasi"
	spaceIndexs := []int{3, 8, 14}

	type args struct {
		palindrome  string
		index       int
		source      string
		spaceIndexs []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				palindrome:  "aku",
				index:       0,
				source:      source,
				spaceIndexs: spaceIndexs,
			},
			want: "aku",
		},
		{
			name: "test 2",
			args: args{
				palindrome:  "suka",
				index:       3,
				source:      source,
				spaceIndexs: spaceIndexs,
			},
			want: "suka",
		},
		{
			name: "test 3",
			args: args{
				palindrome:  "akusuka",
				index:       0,
				source:      source,
				spaceIndexs: spaceIndexs,
			},
			want: "aku suka",
		},
		{
			name: "test 4",
			args: args{
				palindrome:  "sukamakan",
				index:       3,
				source:      source,
				spaceIndexs: spaceIndexs,
			},
			want: "suka makan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSpace(tt.args.palindrome, tt.args.index, tt.args.source, tt.args.spaceIndexs); got != tt.want {
				t.Errorf("addSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_addSpace(t *testing.T) {
// 	type args struct {
// 		text        string
// 		spaceIndexs []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "aku suka makan nasi",
// 			args: args{
// 				text:        "akusukamakannasi",
// 				spaceIndexs: []int{3, 8, 14},
// 			},
// 			want: "aku suka makan nasi",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := addSpace(tt.args.text, tt.args.spaceIndexs); got != tt.want {
// 				t.Errorf("addSpace() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_getFirstLongPalindrome(t *testing.T) {
	type args struct {
		maps map[string]int
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantIndex int
	}{
		{
			name: "longest",
			args: args{maps: map[string]int{
				"ada":   1,
				"kasur": 10,
			}},
			want:      "kasur",
			wantIndex: 10,
		},
		{
			name: "first index",
			args: args{maps: map[string]int{
				"rusak": 15,
				"kasur": 10,
			}},
			want:      "kasur",
			wantIndex: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotIndex := getFirstLongPalindrome(tt.args.maps)
			if got != tt.want {
				t.Errorf("getFirstLongPalindrome() = %v, want %v", got, tt.want)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("getFirstLongPalindrome() gotindex = %v, wantindex %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func Test_getPalindromeMaps(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "not exist",
			args: args{a: "satu", b: "bobo"},
			want: make(map[string]int),
		},
		{
			name: "exist one char",
			args: args{a: "satu", b: "dua"},
			want: map[string]int{
				"a": 1,
				"u": 3,
			},
		},
		{
			name: "exist",
			args: args{a: "adakasurrusak", b: "kasurrusakada"},
			want: map[string]int{
				"a":          0,
				"ada":        0,
				"ak":         11,
				"aka":        2,
				"asurrusak":  4,
				"da":         1,
				"k":          12,
				"ka":         3,
				"kasurrusak": 3,
				"r":          7,
				"rrusak":     7,
				"rusak":      8,
				"s":          5,
				"sak":        10,
				"surrusak":   5,
				"u":          6,
				"urrusak":    6,
				"usak":       9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPalindromeMaps(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPalindromeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPalindrome(t *testing.T) {
	type args struct {
		a string
		b string
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "a",
			args: args{a: "a", b: "a", i: 0, j: 0},
			want: "a",
		},
		{
			name: "empty",
			args: args{a: "a", b: "b", i: 0, j: 0},
			want: "",
		},
		{
			name: "ada",
			args: args{a: "beliada", b: "adaileb", i: 4, j: 0},
			want: "ada",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPalindrome(tt.args.a, tt.args.b, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("getPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_revert(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{text: ""},
			want: "",
		},
		{
			name: "abcde",
			args: args{text: "abcde"},
			want: "edcba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revert(tt.args.text); got != tt.want {
				t.Errorf("revert() = %v, want %v", got, tt.want)
			}
		})
	}
}
