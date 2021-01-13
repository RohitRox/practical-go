package algorithms

import (
	"strings"
	"sort"
	"reflect"
)

type P map[string]int

func AnargramCheck(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	str1Split := strings.Split(strings.ToLower(str1), "")
	str2Split := strings.Split(strings.ToLower(str2), "")

	sort.Strings(str1Split)
	sort.Strings(str2Split)

	return reflect.DeepEqual(str1Split, str2Split)

	// p := P{}

	// for _, c := range str1 {
	// 	s := string(c)

	// 	if _, ok := p[s]; ok {
	// 		p[s] = p[s] + 1
	// 	} else {
	// 		p[s] = 1
	// 	}
	// }

	// for _, c := range str2 {
	// 	s := string(c)

	// 	if k, ok := p[s]; ok {

	// 		if (k - 1 > 0) {
	// 			p[s] = k - 1
	// 		} else {
	// 			delete(p, s)
	// 		}

	// 	} else {
	// 		return false
	// 	}
	// }

	// return len(p) == 0
}
