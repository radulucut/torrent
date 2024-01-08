package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_readDict(t *testing.T) {
	t.Run("should return a map[string]Token", func(t *testing.T) {
		tokens, _ := readDict(0, []byte("d3:cow3:moo4:spam4:eggs4:listl1:al1:bee3:inti-10e3dog4:bark4dictd3foo3:baree"))
		if diff := cmp.Diff(map[string]Token{
			"cow":  [2]int{8, 11},
			"spam": [2]int{19, 23},
			"list": []Token{
				[2]int{32, 33},
				[]Token{
					[2]int{36, 37},
				},
			},
			"int": [2]int{45, 48},
			"dog": [2]int{55, 59},
			"dict": map[string]Token{
				"foo": [2]int{71, 74},
			},
		}, tokens); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	})
}
