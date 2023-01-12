package main

// See https://leetcode.cn/problems/check-if-number-has-equal-digit-count-and-digit-value/ for more details
func digitCount(num string) bool {
	record := make(map[int]int, len(num))
	// 统计每个数字出现的次数
	for i := 0; i < len(num); i++ {
		u := num[i] - '0'
		record[int(u)]++
	}

	// 每个索引出现的次数和对应索引值的比较
	for i := 0; i < len(num); i++ {
		if record[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
