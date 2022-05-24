package main

import (
	"fmt"
)

func main() {
	//grid := [][]int{
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//}

	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	island := maxAreaOfIsland(grid)
	fmt.Println(island)
}

var m4, n4 = 0, 0

func maxAreaOfIsland(grid [][]int) int {
	m4, n4 = len(grid), len(grid[0])

	res := 0
	for i := 0; i < m4; i++ {
		for j := 0; j < n4; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs_04(grid, i, j))
			}
		}
	}
	return res
}

func dfs_04(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= m4 || j >= n4 {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	return dfs_04(grid, i-1, j) +
		dfs_04(grid, i+1, j) +
		dfs_04(grid, i, j-1) +
		dfs_04(grid, i, j+1) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
