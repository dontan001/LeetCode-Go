package leetcode

func ReverseString(s []byte) {

	result := make([]byte, 0)
	reverse(s, 0, &result)
	for i := 0; i < len(result); i++ {
		s[i] = result[i]
	}
}

func reverse(input []byte, begin int, output *[]byte) {
	if begin == len(input) {
		return
	}

	reverse(input, begin+1, output)
	*output = append(*output, input[begin])
}
