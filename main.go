package main

import (
	"fmt"
)

func main() {
	fmt.Println("Running Array Tests")
	target := 0
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, target))
	nums = []int{0, 0, 0, 0}
	fmt.Println(fourSum(nums, target))

	fmt.Println("==================================================")
	fmt.Printf("Running Strings Tests\n")
	fmt.Printf("All Unique: %v\n", unique("hello"))
	fmt.Printf("All Unique: %v\n", unique("helios"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
