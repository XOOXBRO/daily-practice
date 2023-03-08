package main

/*

在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

eg1:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

See https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/description/ for more details
*/

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			dp[i][j] = grid[i][j]
			// 左 上
			var (
				left int
				up   int
			)
			if j > 0 {
				left = dp[i][j-1]
			}
			if i > 0 {
				up = dp[i-1][j]
			}

			var cal = up
			if left > up {
				cal = left
			}

			dp[i][j] += cal
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
