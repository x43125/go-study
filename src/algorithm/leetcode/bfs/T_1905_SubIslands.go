package main

import "fmt"

func main() {
	grid1 := [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}
	grid2 := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}
	islands := countSubIslands(grid1, grid2)
	fmt.Println(islands)
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])

	// first to judge which must is not be a sub is land
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				// then flood fill these islands
				dfs_05(grid2, i, j)
			}
		}
	}

	// last count the other island
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				// then flood fill these islands
				res++
				dfs_05(grid2, i, j)
			}
		}
	}

	return res
}

func dfs_05(grid2 [][]int, i int, j int) {
	m, n := len(grid2), len(grid2[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid2[i][j] == 0 {
		return
	}

	grid2[i][j] = 0

	dfs_05(grid2, i-1, j)
	dfs_05(grid2, i+1, j)
	dfs_05(grid2, i, j-1)
	dfs_05(grid2, i, j+1)
}
