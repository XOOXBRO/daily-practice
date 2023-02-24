package hot100

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
eg1:
输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"

eg2:
输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

See https://leetcode.cn/problems/palindromic-substrings/?favorite=2cktkvj for more details
*/

func countSubstrings(s string) int {
	var (
		record = make([][]bool, len(s))
		res    int
	)
	for i := 0; i < len(s); i++ {
		record[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || record[i+1][j-1]) {
				res++
				record[i][j] = true
			}
		}
	}

	return res
}
