package algorithms

import (
	"fmt"
	"testing"
)

func TestPalindromeCheck(t *testing.T) {
	cases := []struct{
		str string
		isValid bool
	}{
		{ "milim", true },
		{ "milin", false },
		{ "mmmm", true },
		{ "abab", false },
		{ "babazabab", true },
		{ "babaabab", true },
	}

	for _, item := range cases {

		t.Run(fmt.Sprintf("Running %v", item.str), func(t *testing.T){
			res := PalindromeCheck(item.str)

			if res != item.isValid {
				t.Errorf("Expected: %t Got: %t", item.isValid, res)
			}
		})
	}

}

