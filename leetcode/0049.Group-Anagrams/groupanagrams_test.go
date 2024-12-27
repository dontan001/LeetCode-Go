package leetcode

import (
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		word := "eat"
		chars := strings.Split(word, "")
		sort.Strings(chars)
		word = strings.Join(chars, "")
		t.Logf(word)
	})
}
