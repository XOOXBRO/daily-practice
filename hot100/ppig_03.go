package hot100

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

eg1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

eg2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

eg3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

tip:
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

See https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?favorite=2cktkvj for more details
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var max int
	rec := make([]int, len(s))
	rec[0] = 1
	for i := 1; i < len(s); i++ {
		// 记录开始索引
		si := i - rec[i-1]
		for j := si; j < i; j++ {
			if s[j] == s[i] {
				rec[i] = i - j
				break
			}
		}

		if rec[i] == 0 {
			rec[i] = rec[i-1] + 1
		}

		if rec[i] > max {
			max = rec[i]
		}
	}
	return max
}
