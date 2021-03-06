package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1}}

	island := closedIsland(grid)
	fmt.Println(island)
}

var m1, n1 = 0, 0

func closedIsland(grid [][]int) int {
	m1, n1 = len(grid), len(grid[0])
	res := 0

	for i := 0; i < m1; i++ {
		dfs_02(grid, i, 0)
		dfs_02(grid, i, n1-1)
	}

	for j := 0; j < n1; j++ {
		dfs_02(grid, 0, j)
		dfs_02(grid, m1-1, j)
	}

	for i := 0; i < m1; i++ {
		for j := 0; j < n1; j++ {
			if grid[i][j] == 0 {
				res++
				dfs_02(grid, i, j)
			}
		}
	}
	return res
}

func dfs_02(grid [][]int, i int, j int) {
	if i < 0 || j < 0 || i >= m1 || j >= n1 {
		return
	}

	if grid[i][j] == 1 {
		return
	}

	grid[i][j] = 1
	dfs_02(grid, i-1, j)
	dfs_02(grid, i+1, j)
	dfs_02(grid, i, j-1)
	dfs_02(grid, i, j+1)
}
