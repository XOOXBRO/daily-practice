package hot100

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

eg1:
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

eg2:
输入：s = "cbbd"
输出："bb"

See https://leetcode.cn/problems/longest-palindromic-substring/description/?favorite=2cktkvj for more details
*/

func longestPalindrome(s string) string {
	var (
		maxlen int
		left   int
		right  int
		dp     = make([][]bool, len(s))
	)
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}

			if dp[i][j] && j-i+1 > maxlen {
				maxlen = j - i + 1
				left = i
				right = j
			}
		}
	}

	return s[left : right+1]
}
