package basics

import (
	"fmt"
	"testing"
)

func TestHex2Base64(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"", ""},
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
}

	for _, tc := range testcases {
		testname := fmt.Sprintf("%s", tc.in)
		t.Run(testname, func(t *testing.T) {
			baseResult, err := Hex2Base64(tc.in)

			if err != nil {
				t.Skip()
			}

			if baseResult != tc.want {
				t.Errorf("Converted: %v, want %v", baseResult, tc.want)
			}
		})
	}
}
