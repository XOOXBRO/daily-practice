package main

/*

现有一台饮水机，可以制备冷水、温水和热水。每秒钟，可以装满 2 杯 不同 类型的水或者 1 杯任意类型的水。

给你一个下标从 0 开始、长度为 3 的整数数组 amount ，其中 amount[0]、amount[1] 和 amount[2] 分别表示需要装满冷水、温水和热水的杯子数量。返回装满所有杯子所需的 最少 秒数。

See https://leetcode.cn/problems/minimum-amount-of-time-to-fill-cups/ for more details
*/

/*
func fillCups(amount []int) int {
	var res int

	for {
		if amount[0] == amount[1] && amount[1] == amount[2] && amount[0] == 0 {
			break
		}
		if amount[0] == 0 {
			if amount[1] > amount[2] {
				res += amount[1]
			} else {
				res += amount[2]
			}
			break
		}
		if amount[1] == 0 {
			if amount[0] > amount[2] {
				res += amount[0]
			} else {
				res += amount[2]
			}
			break
		}
		if amount[2] == 0 {
			if amount[0] > amount[1] {
				res += amount[0]
			} else {
				res += amount[1]
			}
			break
		}

		var b int
		var s int
		if amount[0] > amount[1] {
			s = 1
		} else {
			b = 1
		}
		if amount[b] < amount[2] {
			b = 2
		}
		if amount[s] > amount[2] {
			s = 2
		}

		amount[b] -= 1
		amount[s] -= 1
		res++
	}
	return res
}
*/

/*
想法:
	每次总是用 最大值 - 最小值
注意点：
	规避掉最小值为 0 的情况
*/
