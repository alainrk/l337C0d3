package l337C0d3

import (
	"reflect"
	"testing"
)

func Test_lcHelper(t *testing.T) {
	type args struct {
		digits   string
		cache    map[string][]string
		solution map[string]bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := lcHelper(tt.args.digits, tt.args.cache, tt.args.solution); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lcHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
