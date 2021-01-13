package algorithms

import "fmt"

func merge(a []int, b []int) []int {

	size, i, j := len(a)+len(b), 0, 0
	result := make([]int, size)

	for k := 0; k < size; k++ {

		switch {
		case i == len(a):
			result[k] = b[j]
			j += 1
		case j == len(b):
			result[k] = a[i]
			i += 1
		case a[i] > b[j]:
			result[k] = b[j]
			j += 1
		case a[i] < b[j]:
			result[k] = a[i]
			i += 1
		default:
			result[k] = b[j]
			j += 1
		}

	}

	return result

}

func mSort(arr []int) []int {
	fmt.Println(arr)
	if len(arr) < 2 {
		return arr
	} else {
		l := len(arr)
		p := l / 2

		left := arr[0:p]
		right := arr[p:l]

		return merge(mSort(left), mSort(right))
	}

}

func MergeSort(list []int) []int {
	return mSort(list)
}
