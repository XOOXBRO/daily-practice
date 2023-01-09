package main

func minOperations(nums []int, x int) int {
	if nums[0] > x && nums[len(nums)-1] > x {
		return -1
	}
	if nums[0] == x || nums[len(nums)-1] == x {
		return 1
	}

	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum < x {
		return -1
	}
	if sum == x {
		return len(nums)
	}

	target := sum - x
	return getSubArr(nums, target)
}

// template of sliding window
func getSubArr(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := 0
	sum := 0
	res := -1

	for right < length {
		sum += nums[right]

		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			res = max(res, right-left+1)
		}
		right++
	}

	if res == -1 {
		return -1
	}

	return length - res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
