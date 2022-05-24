package main

import "fmt"

func main() {
	i := combine(4, 3)
	fmt.Println(i)
}

func combine(n int, k int) [][]int {
	//if n2 <= 0 || k <= 0 || k > n2 {
	//	return [][]int{}
	//}

	var res [][]int
	var track []int
	backTrack_01(n, k, 1, track, &res)
	return res
}

func backTrack_01(n int, k int, start int, track []int, res *[][]int) {
	if len(track) == k {
		b := make([]int, len(track))
		copy(b, track)
		*res = append(*res, b)
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backTrack_01(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
