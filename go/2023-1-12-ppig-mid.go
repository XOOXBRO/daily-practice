package main

// Plan1: mock ,but exist string.split,it's slow.
// Time feat 9.52%. Memory feat 19.5%.
func evaluate(s string, knowledge [][]string) string {
	// 将 knowledge 转为 map 存储
	kr := make(map[string]string, len(knowledge))
	for _, v := range knowledge {
		kr[v[0]] = v[1]
	}

	var left int
	var flag bool
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			flag = true
			left = i
			continue
		}
		if !flag {
			continue
		}
		if s[i] == ')' {
			key := s[left+1 : i]
			val, ok := kr[key]
			var next string
			if i+1 < len(s) {
				next = s[i+1:]
			}
			if ok {
				s = s[:left] + val + next
				i = left + len(val) - 1
			} else {
				s = s[:left] + "?" + next
				i = left
			}
			flag = false
		}
	}

	return s
}

// Plan2, use res to save string,so it decrease generate of string,it's more fast.
// Time feat 47.62%. Memory feat 100%.
func evaluate2(s string, knowledge [][]string) string {
	// 将 knowledge 转为 map 存储
	kr := make(map[string]string, len(knowledge))
	for _, v := range knowledge {
		kr[v[0]] = v[1]
	}

	var res []byte
	var left int
	var flag bool
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			flag = true
			left = i
			continue
		}
		if !flag {
			res = append(res, s[i])
			continue
		}
		if s[i] == ')' {
			val, ok := kr[s[left+1:i]]
			if ok {
				res = append(res, []byte(val)...)
			} else {
				res = append(res, '?')
			}
			flag = false
		}
	}

	return string(res)
}
