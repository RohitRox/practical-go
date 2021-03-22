package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInplaceArrayUniq(t *testing.T) {
	cases := []struct {
		data []int
		uniq []int
	}{
		{[]int{1, 2, 2, 3, 4, 5, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1}, []int{1}},
	}

	for _, item := range cases {
		t.Run(fmt.Sprintf("Test Case %v", item.data), func(t *testing.T) {

			got := InplaceArrayUniq(item.data)

			if !reflect.DeepEqual(item.uniq, got) {
				t.Errorf("Expected: %v Got: %v", item.uniq, got)
			}

		})
	}
}
