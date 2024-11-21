package leetcode

func Subsets(nums []int) [][]int {

	var subsetsAll [][]int = make([][]int, 0)
	if len(nums) == 0 {
		subsetsAll = append(subsetsAll, []int{})
		return subsetsAll
	}

	var subsetsNolast = Subsets(nums[:len(nums)-1])
	subsetsAll = append(subsetsAll, subsetsNolast...)

	for _, arr := range subsetsNolast {
		subset := []int{nums[len(nums)-1]}
		subset = append(subset, arr...)
		subsetsAll = append(subsetsAll, subset)
	}

	return subsetsAll
}
