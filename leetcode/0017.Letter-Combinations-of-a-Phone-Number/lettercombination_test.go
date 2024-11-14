package leetcode

import "testing"

func TestLetterCombinations(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		digits := "2"
		combinations := LetterCombinations(digits)
		t.Logf("all combinations: %+v", combinations)
	})

	t.Run("case 2", func(t *testing.T) {
		digits := "23"
		combinations := LetterCombinations(digits)
		t.Logf("all combinations: %+v", combinations)
	})

	t.Run("case 4", func(t *testing.T) {
		digits := "2347"
		combinations := LetterCombinations(digits)
		t.Logf("all combinations: %+v", combinations)
	})

	t.Run("case 5", func(t *testing.T) {
		digits := "23"
		combinations := LetterCombinationsWithLoop(digits)
		t.Logf("all combinations: %+v", combinations)
	})

	t.Run("case 6", func(t *testing.T) {
		digits := "2347"
		combinations := LetterCombinationsWithLoop(digits)
		t.Logf("all combinations: %+v", combinations)
	})
}
