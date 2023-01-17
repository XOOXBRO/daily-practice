package main

// See https://leetcode.cn/problems/count-nice-pairs-in-an-array/description/ for more details
func countNicePairs(nums []int) int {
	var res int
	record := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := nums[i] - rev(nums[i])
		res += record[n]
		record[n]++
	}
	return res % (1e9 + 7)
}

func rev(n int) int {
	var res int
	for n > 0 {
		i := n % 10
		n /= 10
		res = res*10 + i
	}
	return res
}

// [13,10,35,24,76]
// 13-31=-18
// 10-1-=9
// 35-53=-18
// 24-42=-18
// 76-67=9
