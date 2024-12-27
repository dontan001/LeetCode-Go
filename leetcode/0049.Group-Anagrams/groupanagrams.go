package leetcode

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {

	root := func(anagram string) string {
		chars := strings.Split(anagram, "")
		sort.Strings(chars)
		return strings.Join(chars, "")
	}

	groups := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		key := root(strs[i])
		if _, ok := groups[key]; ok {
			groups[key] = append(groups[key], strs[i])
		} else {
			groups[key] = []string{strs[i]}
		}
	}

	result := make([][]string, 0)
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
