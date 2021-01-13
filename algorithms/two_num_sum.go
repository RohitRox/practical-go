package algorithms

type SumHash map[int]int

func TwoNumSum(list []int, sum int) [][]int {

	hash := SumHash{}
	resultList := [][]int{}

	for _, item := range list {

		required := sum - item

		if _, ok := hash[required]; ok {
			resultList = append([][]int{ []int{required, item} }, resultList...)
		}

		if _, ok := hash[item]; !ok {
			hash[item] = item
		}
	}

	return resultList
}
