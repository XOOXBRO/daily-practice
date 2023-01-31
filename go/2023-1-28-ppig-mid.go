package main

// [2,1,6,4]
// [4,1,1,2,5,1,5,4]
// See https://leetcode.cn/problems/ways-to-make-a-fair-array/description/ for more details

// Brutal
func waysToMakeFair(nums []int) int {
	var count int

	for i := 0; i < len(nums); i++ {
		var (
			odd  int
			even int
			flag bool
			pre  = 0
			next = i + 1
		)

		if (i & 1) != 1 {
			flag = true
		}
		for pre < i {
			if (pre & 1) == 1 {
				odd += nums[pre]
			} else {
				even += nums[pre]
			}
			pre++
		}
		for next < len(nums) {
			if !flag {
				odd += nums[next]
			} else {
				even += nums[next]
			}
			flag = !flag
			next++
		}
		if odd == even {
			count++
		}
	}

	return count
}
