package main

import (
	"fmt"
	"sort"
)

func main() {
	//candidates := []int{2, 3, 6, 7}
	candidates := []int{2, 3, 5}
	target := 8
	sum_04 := combinationSum_04(candidates, target)
	fmt.Println(sum_04)
}

var res_04 = [][]int{}
var track_04 = []int{}
var targetV_04 = 0

func combinationSum_04(candidates []int, target int) [][]int {
	res_04 = [][]int{}
	track_04 = []int{}
	targetV_04 = target

	sort.Ints(candidates)

	backTrack_04(candidates, 0, 0)
	return res_04
}

func backTrack_04(candidates []int, sum int, start int) {
	if sum == targetV_04 {
		c := make([]int, len(track_04))
		copy(c, track_04)
		res_04 = append(res_04, c)
	}

	if sum > targetV_04 {
		return
	}

	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		track_04 = append(track_04, candidates[i])
		backTrack_04(candidates, sum, i)
		sum -= candidates[i]
		track_04 = track_04[:len(track_04)-1]
	}
}
