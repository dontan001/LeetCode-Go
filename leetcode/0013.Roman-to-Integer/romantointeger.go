package leetcode

func RomanToInteger(s string) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var index, num int = 0, 0
	for len(s) > 0 {
		symbolLength := len(symbols[index])
		if symbolLength > len(s) {
			index++
			continue
		}

		romanNumeral := s[0:symbolLength]
		if romanNumeral == symbols[index] {
			num = num + values[index]
			s = s[symbolLength:len(s)]
		} else {
			index++
		}
	}

	return num
}
