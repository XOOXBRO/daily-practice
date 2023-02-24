package hot100

/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

eg1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]

eg2:
输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]

eg3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

See https://leetcode.cn/problems/daily-temperatures/?favorite=2cktkvj for more details
*/

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		var (
			curr = temperatures[i]
			idx  = i + 1
			pre  = temperatures[idx]
		)

		for curr >= pre && answer[idx] != 0 {
			idx += answer[idx]
			pre = temperatures[idx]
		}

		if curr < pre {
			answer[i] = idx - i
		}
	}

	return answer
}
