package leetcode

import "testing"

func TestIntegerToRoman(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		var num int = 3749
		romanNumerals := IntegerToRoman(num) //expected MMMDCCXLIX
		t.Logf("the roman numerals of num %d is %s", num, romanNumerals)
	})
}
