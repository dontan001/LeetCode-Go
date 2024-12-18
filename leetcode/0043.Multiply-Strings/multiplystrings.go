package leetcode

import "strconv"

func Multiply(num1 string, num2 string) string {

	size := len(num1) + len(num2)
	product := make([]int, size)
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			h, l := i+j, i+j+1
			temp := int(num1[i]-'0')*int(num2[j]-'0') + product[l]
			product[l] = temp % 10
			product[h] = product[h] + temp/10
		}
	}

	var result string = ""
	for i := 0; i < size; i++ {
		if product[i] == 0 && len(result) == 0 {
			continue
		} else {
			result += strconv.Itoa(product[i])
		}
	}

	if result == "" {
		return "0"
	}

	return result
}
