package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

var res = [][]int{}
var track = []int{}

func subsetsWithDup(nums []int) [][]int {
	res = [][]int{}
	track = []int{}
	sort.Ints(nums)
	backTrack_02(nums, 0)
	return res
}

func backTrack_02(nums []int, start int) {
	c := make([]int, len(track))
	copy(c, track)
	res = append(res, c)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backTrack_02(nums, i+1)
		track = track[:len(track)-1]
	}
}
