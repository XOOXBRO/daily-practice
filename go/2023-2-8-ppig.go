package main

import (
	"sort"
	"strings"
)

/*
你是一位系统管理员，手里有一份文件夹列表 folder，你的任务是要删除该列表中的所有 子文件夹，并以 任意顺序 返回剩下的文件夹。

如果文件夹 folder[i] 位于另一个文件夹 folder[j] 下，那么 folder[i] 就是 folder[j] 的 子文件夹 。

文件夹的「路径」是由一个或多个按以下格式串联形成的字符串：'/' 后跟一个或者多个小写英文字母。

例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路径，而空字符串和 "/" 不是。

*/

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	var res []string
	res = append(res, folder[0])

	for i := 1; i < len(folder); i++ {
		pre := res[len(res)-1]
		if !strings.HasPrefix(folder[i], pre) || folder[i][len(pre)] != '/' {
			res = append(res, folder[i])
		}
	}
	return res
}
