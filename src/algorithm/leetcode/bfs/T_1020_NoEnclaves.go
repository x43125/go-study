package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}

	island := numEnclaves(grid)
	fmt.Println(island)
}

var m2, n2 = 0, 0

func numEnclaves(grid [][]int) int {
	m2, n2 = len(grid), len(grid[0])
	res := 0

	for i := 0; i < m2; i++ {
		dfs_03(grid, i, 0)
		dfs_03(grid, i, n2-1)
	}

	for j := 0; j < n2; j++ {
		dfs_03(grid, 0, j)
		dfs_03(grid, m2-1, j)
	}

	for i := 0; i < m2; i++ {
		for j := 0; j < n2; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs_03(grid [][]int, i int, j int) {
	if i < 0 || j < 0 || i >= m2 || j >= n2 {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	dfs_03(grid, i-1, j)
	dfs_03(grid, i+1, j)
	dfs_03(grid, i, j-1)
	dfs_03(grid, i, j+1)
}
