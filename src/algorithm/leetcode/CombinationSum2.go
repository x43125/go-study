package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

var trackSum = 0
var targetV = 0
var res_03 = [][]int{}

func combinationSum2(candidates []int, target int) [][]int {
	res_03 = [][]int{}
	targetV = target

	sort.Ints(candidates)
	track := []int{}
	backTrack_03(candidates, track, 0)
	return res_03
}

func backTrack_03(nums []int, track []int, start int) {
	if trackSum == targetV {
		c := make([]int, len(track))
		copy(c, track)
		res_03 = append(res_03, c)
	}

	if trackSum > targetV {
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		trackSum += nums[i]
		track = append(track, nums[i])
		backTrack_03(nums, track, i+1)
		track = track[:len(track)-1]
		trackSum -= nums[i]
	}
}
