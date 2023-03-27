package main

/*
在一个有向图中，节点分别标记为 0, 1, ..., n-1。图中每条边为红色或者蓝色，且存在自环或平行边。

red_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的红色有向边。类似地，blue_edges 中的每一个 [i, j] 对表示从节点 i 到节点 j 的蓝色有向边。

返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么 answer[x] = -1。

eg:
输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
输出：[0,1,-1]

输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
输出：[0,1,-1]

输入：n = 3, red_edges = [[1,0]], blue_edges = [[2,1]]
输出：[0,-1,-1]

输入：n = 3, red_edges = [[0,1]], blue_edges = [[1,2]]
输出：[0,1,2]

输入：n = 3, red_edges = [[0,1],[0,2]], blue_edges = [[1,0]]
输出：[0,1,1]
See https://leetcode.cn/problems/shortest-path-with-alternating-colors/ for more details
*/

/*
func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	// record 用于构造树
	record := make([][]pair, n)
	for _, v := range redEdges {
		record[v[0]] = append(record[v[0]], pair{x: v[1], color: 0})
	}
	for _, v := range blueEdges {
		record[v[0]] = append(record[v[0]], pair{x: v[1], color: 1})
	}

	answer := make([]int, n)
	for i := 1; i < n; i++ {
		answer[i] = -1
	}

	// 可能会引起死循环，因此已经访问过的必须进行标记
	// todo 标记队列中已经使用过的元素
	p := []pair{{x: 0, color: 0}, {x: 0, color: 1}}
	for level := 0; p != nil; level++ {
		tmp := p
		p = nil
		for _, v := range tmp {
			if answer[v.x] < 0 {
				answer[v.x] = level
			}
			for _, node := range record[v.x] {
				if node.color != v.color {
					p = append(p, node)
				}
			}
		}
	}

	return answer
}

*/

/*
n = 3
redEdges = [[0,1],[0,2]]
blueEdges = [[1,0]]
*/
