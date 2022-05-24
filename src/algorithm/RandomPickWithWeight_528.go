package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	preSum []int
}

// Constructor generate a prefix sum array
func Constructor(w []int) Solution {
	preSum := make([]int, len(w))
	preSum[0] = w[0]

	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i-1] + w[i]
	}

	return Solution{preSum: preSum}
}

func (this *Solution) PickIndex() int {
	length := len(this.preSum)-1
	n := rand.Intn(this.preSum[length]) + 1

	left, right := 0, length
	for left < right {
		mid := (right-left)>>1 + left
		if this.preSum[mid] < n {
			left = mid + 1
		} else {
			// 其他情况都将right=mid 来缩小查询范围，以找到最左侧的第一个大于等于他的值
			right = mid
		}
	}

	return left
}

func main() {
	w := []int{1, 2, 3, 4, 5}
	solution := Constructor(w)
	res := solution.PickIndex()
	fmt.Println(res)
}
