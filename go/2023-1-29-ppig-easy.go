package main

// See https://leetcode.cn/problems/count-asterisks/description/ for more details
func countAsterisks(s string) int {
	var (
		count int
		res   int
	)

	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			count++
		}
		if s[i] == '*' && (count&1) == 0 {
			res++
		}
	}

	return res
}

/*
题解
题目的意思是 每两个'|' '|' 为一对，要找到不在一对 '|' 中 ‘*’ 的数量。
	第一个 '|' 只能和第二个凑成一对，第三个只能和第四个凑成一对
	第二个 '|' 和第三个 ‘|’ 不能凑成一对
	| 的总数量恰好为偶数
	因此，只要 | 的数量为偶数时出现的 * 绝不会在任意一对 ‘|’之间。
	倘若 | 的数量为奇数，说明 必然有一个 | 对应，因此其中出现的 *，在 ‘|’ 对中。
*/
