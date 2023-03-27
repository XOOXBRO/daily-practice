package main

/*
给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：

从 nums 选出 两个 相等的 整数
从 nums 中移除这两个整数，形成一个 数对
请你在 nums 上多次执行此操作直到无法继续执行。

返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，其中 answer[0] 是形成的数对数目，answer[1] 是对 nums 尽可能执行上述操作后剩下的整数数目。

See https://leetcode.cn/problems/maximum-number-of-pairs-in-array/ for more details
*/

/*
func numberOfPairs(nums []int) []int {
	answer := make([]int, 2)
	record := make(map[int]int)

	for _, v := range nums {
		record[v] += 1
		if record[v] == 2 {
			record[v] = 0
			answer[0] += 1
		}
	}

	for _, v := range record {
		if v == 1 {
			answer[1] += 1
		}
	}
	return answer
}
*/
