package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	islands := numIslands(grid)
	fmt.Println(islands)
}

var m, n = 0, 0

func numIslands(grid [][]byte) int {
	m, n = len(grid), len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
