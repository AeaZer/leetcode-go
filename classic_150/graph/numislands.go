package graph

// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/

/*给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。*/

const (
	spaceWater byte = iota
	spaceLand
)

func numIslands(grid [][]byte) int {
	var res int
	iMax, jMax := len(grid), len(grid[0])

	return res
}
