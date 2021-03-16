package algorithms

import (
	"fmt"
	"testing"
)

func TestAnargramCheck(t *testing.T) {
	cases := []struct {
		str1    string
		str2    string
		isValid bool
	}{
		{"fired", "fried", true},
		{"gainly", "laying", true},
		{"xman", "xMen", false},
		{"sadder", "dreads", true},
		{"panic", "panik", false},
		{"knee", "keen", true},
		{"meeren", "meren", false},
		{"Listen", "slient", true},
	}

	for _, item := range cases {
		t.Run(fmt.Sprintf("Running %v - %v", item.str1, item.str2), func(t *testing.T) {
			res := AnargramCheck(item.str1, item.str2)

			if res != item.isValid {
				t.Errorf("Expected: %t Got: %t", item.isValid, res)
			}
		})
	}
}
