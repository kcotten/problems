package main

import (
	"fmt"
)

func main() {
	fmt.Println("Running Test")
	target := 0
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, target))
	nums = []int{0, 0, 0, 0}
	fmt.Println(fourSum(nums, target))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
