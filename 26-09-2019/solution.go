package solution

func Solution(input []int, target int) [2]int {
	left := search(input, target, 0, len(input)-1, true)
	if left == -1 {
		return [2]int{-1, -1}
	}
	right := search(input, target, left, len(input)-1, false)

	return [2]int{left, right}

}

func search(nums []int, target, low, high int, isLeft bool) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if nums[mid] > target || (nums[mid] == target && isLeft && nums[mid] == nums[mid-1]) {
		return search(nums, target, low, mid-1, isLeft)
	} else if nums[mid] < target || (nums[mid] == target && !isLeft && nums[mid] == nums[mid+1]) {
		return search(nums, target, mid+1, high, isLeft)
	}

	return mid

}
