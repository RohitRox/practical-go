package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {

	cases := []struct {
		data   []int
		sorted []int
	}{
		{[]int{4, 1}, []int{1, 4}},
		{[]int{4, 1, 5}, []int{1, 4, 5}},
		{[]int{4, 1, 1}, []int{1, 1, 4}},
		{[]int{1, 3, 5, 2, 4}, []int{1, 2, 3, 4, 5}},
	}

	for _, item := range cases {

		t.Run(fmt.Sprintf("Sort %v", item.data), func(t *testing.T) {

			sorted := MergeSort(item.data)
			if !reflect.DeepEqual(sorted, item.sorted) {
				t.Errorf("Expected: %v Got: %v", item.sorted, sorted)
			}

		})

	}
}
