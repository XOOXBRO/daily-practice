package main

import "sort"

/*
给你两个二维整数数组 items1 和 items2 ，表示两个物品集合。每个数组 items 有以下特质：

items[i] = [valuei, weighti] 其中 valuei 表示第 i 件物品的 价值 ，weighti 表示第 i 件物品的 重量 。
items 中每件物品的价值都是 唯一的 。
请你返回一个二维数组 ret，其中 ret[i] = [valuei, weighti]， weighti 是所有价值为 valuei 物品的 重量之和 。

注意：ret 应该按价值 升序 排序后返回。

See https://leetcode.cn/problems/merge-similar-items/ for more details
*/

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})
	sort.Slice(items2, func(i, j int) bool {
		return items2[i][0] < items2[j][0]
	})

	var (
		i1  int
		i2  int
		res [][]int
	)

	for i1 < len(items1) && i2 < len(items2) {
		if items1[i1][0] == items2[i2][0] {
			res = append(res, []int{items1[i1][0], items1[i1][1] + items2[i2][1]})
			i1++
			i2++
			continue
		}
		if items1[i1][0] < items2[i2][0] {
			res = append(res, items1[i1])
			i1++
		} else {
			res = append(res, items2[i2])
			i2++
		}
	}

	for ; i1 < len(items1); i1++ {
		res = append(res, items1[i1])
	}
	for ; i2 < len(items2); i2++ {
		res = append(res, items2[i2])
	}
	return res
}
