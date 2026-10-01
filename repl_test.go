package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct{
		input string
		want []string
	}{
		{input: "hello world", want: []string{"hello", "world"}},
		{input: "     HEllo  World", want: []string{"hello", "world"}},
		{input: "     hello   ", want: []string{"hello"}},
		{input: "", want: []string{}},
	}

	for _, tc := range tests {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}