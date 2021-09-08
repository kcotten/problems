package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
