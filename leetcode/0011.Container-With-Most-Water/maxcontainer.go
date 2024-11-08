package leetcode

func MaxArea(height []int) int {

	var max int = 0
	var left, right int = 0, len(height) - 1
	for left < right {
		h := height[left]
		if h > height[right] {
			h = height[right]
		}

		area := h * (right - left)
		if area > max {
			max = area
		}

		if h == height[left] {
			left++
		} else {
			right--
		}
	}

	return max
}
