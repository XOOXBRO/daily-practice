package main

// See https://leetcode.cn/problems/longest-palindromic-substring/description/ for more details
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
