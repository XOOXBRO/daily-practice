package main

import (
	"sort"
)

/*
给你一个非负整数数组 nums 。在一步操作中，你必须：

* 选出一个正整数 x ，x 需要小于或等于 nums 中 最小 的 非零 元素。
* nums 中的每个正整数都减去 x。
返回使 nums 中所有元素都等于 0 需要的 最少 操作数。

eg01:
输入：nums = [1,5,0,3,5]
输出：3
解释：
第一步操作：选出 x = 1 ，之后 nums = [0,4,0,2,4] 。
第二步操作：选出 x = 2 ，之后 nums = [0,2,0,0,2] 。
第三步操作：选出 x = 2 ，之后 nums = [0,0,0,0,0] 。

eg02:
输入：nums = [0]
输出：0
解释：nums 中的每个元素都已经是 0 ，所以不需要执行任何操作。

See https://leetcode.cn/problems/make-array-zero-by-subtracting-equal-amounts/ for more details
*/
func minimumOperations(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 0
		}
		return 1
	}

	sort.Ints(nums)
	var res int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != 0 && nums[i] != nums[i+1] {
			res++
		}
	}

	if nums[0] != 0 || res != 0 {
		return res + 1
	}
	return res
}
