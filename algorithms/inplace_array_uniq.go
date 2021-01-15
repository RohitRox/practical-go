package algorithms

func InplaceArrayUniq(list []int) []int {
	j := 0
	for i := j + 1; i < len(list); i++ {
		if list[i] != list[i-1] {
			j++
			list[j] = list[i]
		}
	}
	return list[:j+1]
}
