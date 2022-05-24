package main

import "fmt"

/**
binary search
*/
func main() {
	nums := []int{1, 2, 2, 2, 5}
	target := 1
	search := binarySearch(nums, target)
	fmt.Println(search)

	traverse := binarySearchTraverse(nums, target, 0, len(nums)-1)
	fmt.Println(traverse)

	left := BinarySearchLeft(nums, target)
	fmt.Println(left)
}

/**
  func_1: iteration
*/
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high-low)>>1 + low

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

/**
  func_2: traverse
*/
func binarySearchTraverse(nums []int, target int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (high-low)>>1 + low
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearchTraverse(nums, target, mid+1, high)
	} else {
		return binarySearchTraverse(nums, target, low, mid-1)
	}
}

func BinarySearchLeft(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)>>1 + left
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	}
	return -1
}
