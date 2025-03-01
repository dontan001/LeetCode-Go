package leetcode

var DigitToLetters map[string][]string = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	return combine(digits)
}

func LetterCombinationsWithLoop(digits string) []string {
	allCombinations := []string{}

	for len(digits) > 0 {
		combinations := []string{}

		letters := DigitToLetters[string(digits[0])]
		for _, letter := range letters {
			if len(allCombinations) == 0 {
				combinations = append(combinations, letter)
				continue
			}

			for _, combination := range allCombinations {
				combinations = append(combinations, combination+letter)
			}
		}

		allCombinations = combinations
		digits = digits[1:]
	}

	return allCombinations
}

func combine(digits string) []string {
	letters := DigitToLetters[string(digits[0])]
	if len(digits) == 1 {
		return letters
	}

	combinationsAll := []string{}
	combinations := combine(digits[1:len(digits)])
	for _, letter := range letters {
		for _, combination := range combinations {
			combinationsAll = append(combinationsAll, letter+combination)
		}
	}
	return combinationsAll
}
