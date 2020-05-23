package problem0165

import (
	"strconv"
	"strings"
)

func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}

	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	for i := 0; i < len(v1); i++ {
		vs1, _ := strconv.Atoi(v1[i])
		vs2, _ := strconv.Atoi(v2[i])
		if vs1 == vs2 {
			continue
		} else if vs1 > vs2 {
			return 1
		} else {
			return -1
		}
	}

	return 0
}
