package main

// See https://leetcode.cn/problems/rearrange-characters-to-make-target-string/ for more details
func rearrangeCharacters(s string, target string) int {
	if len(s) < len(target) {
		return 0
	}

	record := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}

	var count int
	var flag bool
	for !flag {
		for i := 0; i < len(target); i++ {
			record[target[i]-'a']--
			if record[target[i]-'a'] < 0 {
				flag = true
				break
			}
		}
		if !flag {
			count++
		}
	}

	return count
}

/*
阅读题干
意思就是从 s 中找到能组成 target 的次数且 s 中的字符串不能重复使用。
不能重复使用的意思是 s 中倘若部分字符已经组成了 target ，那么这些字符就不能再次用来组成 target。

解题思路
第一想法就是用 哈希表 把 target 存起来，然后在 s 里面判断字符出现的字数，但是问题就来了：
	怎么判断字符串 s 能够组成 target？ 倘若 哈希表 里面存储的是字符对应的次数，那么如果在 s 还没遍历完时 哈希表里面字符的次数都归 0 了，
	便意味着可以组成一次 target，但是后面还有没遍历完的字符该怎么办呢？ 那么必然涉及到 哈希表 的复制，重复使用，空间复杂度很高。

于是反过来想，用 哈希表存储 s ，然后不断重复遍历 target，直到 s 里对应的字符出现次数为0。
时间复杂度，O(len(s))
*/
