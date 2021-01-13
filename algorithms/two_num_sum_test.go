package algorithms

import (
	"fmt"
	"testing"
	"reflect"
)

func TestTwoNumSum(t *testing.T) {
	cases := []struct {
		data []int
		sum int
		results [][]int
	} {
		{ []int{1,2,3,4,5}, 6, [][]int{ []int{1,5}, []int{2,4} }},
		{ []int{2,4,6,8}, 8, [][]int{ []int{2,6} }},
		{ []int{5, 10, 15, 20, 25, 30}, 16, [][]int{ }},
	}

	for _, item := range cases {

		t.Run(fmt.Sprintf("Runing case %v", item.data), func(t *testing.T) {
			results := TwoNumSum(item.data, item.sum)

			if !reflect.DeepEqual(results, item.results) {
				t.Errorf("Failed for case %v Expected: %v, Got: %v", item.data, item.results, results)
			}
		})

	}

}
