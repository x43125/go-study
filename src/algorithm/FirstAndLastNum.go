package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	res := []int{-1, -1}
	for left <= right {
		mid := (right-left)>>1 + left
		midV := nums[mid]
		if midV < target {
			left = mid + 1
		} else if midV > target {
			right = mid - 1
		} else {
			res[0] = searchLeft(nums, target, left, mid)
			res[1] = searchRight(nums, target, mid, right)
			break
		}
	}

	return res
}

func searchLeft(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := (right-left)>>1 + left
		midV := nums[mid]
		if midV < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchRight(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := (right-left)>>1 + left
		midV := nums[mid]
		if midV > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums := []int{1, 2, 2, 2, 2, 3, 3, 4}
	res := searchRange(nums, 3)
	fmt.Println(res[0], res[1])
}
