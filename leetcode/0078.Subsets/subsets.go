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

func SubsetMask(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	subsetsAll := make([][]int, 0)
	size := len(nums)
	for i := 0; i < (1 << size); i++ {
		subset := []int{}
		for j := 0; j < size; j++ {
			include := i & (1 << j)
			if include > 0 {
				subset = append(subset, nums[size-j-1])
			}
		}

		subsetsAll = append(subsetsAll, subset)
	}

	return subsetsAll
}
