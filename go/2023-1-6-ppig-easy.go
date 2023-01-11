package main

// See https://leetcode.cn/problems/count-integers-with-even-digit-sum/ for details
func countEven(num int) int {
	var res int
	for i := 1; i <= num; i++ {
		j := i
		sum := 0
		for j >= 10 {
			sum += j % 10
			j /= 10
		}
		sum += j

		if (sum & 1) != 1 {
			res++
		}
	}

	return res
}
