package algorithms

type SumHash map[int]int
type RL [][]int

func TwoNumSum(list []int, sum int) [][]int {

	hash := SumHash{}
	resultList := RL{}

	for _, item := range list {

		required := sum - item

		if _, ok := hash[required]; ok {
			resultList = append(RL{[]int{required, item}}, resultList...)
		}

		if _, ok := hash[item]; !ok {
			hash[item] = item
		}
	}

	return resultList
}
