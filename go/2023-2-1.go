package main

/*
给你字符串 key 和 message ，分别表示一个加密密钥和一段加密消息。解密 message 的步骤如下：

使用 key 中 26 个英文小写字母第一次出现的顺序作为替换表中的字母 顺序 。
将替换表与普通英文字母表对齐，形成对照表。
按照对照表 替换 message 中的每个字母。
空格 ' ' 保持不变。
例如，key = "happy boy"（实际的加密密钥会包含字母表中每个字母 至少一次），据此，可以得到部分对照表（'h' -> 'a'、'a' -> 'b'、'p' -> 'c'、'y' -> 'd'、'b' -> 'e'、'o' -> 'f'）。
返回解密后的消息。

See https://leetcode.cn/problems/decode-the-message/description/ for more details
*/
/*
func decodeMessage(key string, message string) string {
	en := byte('a')
	record := map[rune]byte{}

	for _, c := range key {
		if c == ' ' {
			continue
		}
		if _, ok := record[c]; !ok {
			record[c] = en
			en++
		}
	}

	res := []byte(message)
	for i, v := range res {
		if v != ' ' {
			res[i] = record[rune(v)]
		}
	}
	return string(res)
}
*/
