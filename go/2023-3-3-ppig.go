package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给你一个长度为 n 的字符串数组 names 。你将会在文件系统中创建 n 个文件夹：在第 i 分钟，新建名为 names[i] 的文件夹。

由于两个文件 不能 共享相同的文件名，因此如果新建文件夹使用的文件名已经被占用，系统会以 (k) 的形式为新文件夹的文件名添加后缀，其中 k 是能保证文件名唯一的 最小正整数 。

返回长度为 n 的字符串数组，其中 ans[i] 是创建第 i 个文件夹时系统分配给该文件夹的实际名称。

See https://leetcode.cn/problems/making-file-names-unique/ for more details
*/
func getFolderNames(names []string) []string {
	var (
		res    = make([]string, 0, len(names))
		rec    = make(map[string]int, len(names))
		length = len(names)
	)

	for i := 0; i < length; i++ {
		cs := names[i]
		count, ok := rec[cs]
		if ok {
			cs += fmt.Sprintf("%s%d%s", "(", count, ")")
			var c int
			c, ok = rec[cs]
			count = c
			for ok {
				index := strings.LastIndex(cs, "(")
				u := cs[index+1 : index+2]
				n, err := strconv.Atoi(u)
				if err != nil {
					return nil
				}
				count += n
				cs = fmt.Sprintf("%s%d%s", cs[:index+1], count, cs[index+2:])
				count, ok = rec[cs]
			}
		}
		rec[cs] = count
		if count == 0 {
			rec[cs] += 1
		}
		res = append(res, cs)
	}
	return res
}

func main() {
	//res := getFolderNames([]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"})
	res := getFolderNames([]string{"wano", "wano", "wano", "wano"})
	fmt.Println(res)
}
