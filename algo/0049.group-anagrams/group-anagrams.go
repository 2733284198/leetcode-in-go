package problem0049

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string

	if len(strs) == 0 {
		return result
	}

	helpMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		var curStrArr = strings.Split(strs[i], "")

		sort.Strings(curStrArr)

		var curStrIndex = strings.Join(curStrArr, "")

		helpMap[curStrIndex] = append(helpMap[curStrIndex], strs[i])
	}

	for _, value := range helpMap {
		result = append(result, value)
	}

	return result
}
