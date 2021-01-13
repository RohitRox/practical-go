package algorithms

func PalindromeCheck(str string) bool {

	l := len(str)

	mid := l/2

	for i,j := 0, l-1; i < mid || j > mid; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
