package leetcode

func IntegerToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var romanNumerals string
	var index int = 0
	for num > 0 {
		if num-values[index] >= 0 {
			romanNumerals = romanNumerals + symbols[index]
			num = num - values[index]
		} else {
			index++
		}
	}

	return romanNumerals
}
