package main

import "math"

// Start by defining the max and min integers -> the node cannot contain a num bigger or smaller
func isValidBST(root *TreeNode) bool {
	MAX := int(^uint(0) >> 1)
	MIN := ^MAX
	return valid(root, MIN, MAX)
}

//
func valid(node *TreeNode, low int, high int) bool {
	// Base case we made it to a leaf -> tree is ok at this leaf
	if node == nil {
		return true
	}
	// Val will either be larger or smaller than the preceding number from the recursive call
	// if not then return false -> short circuiting the next return statement
	if node.Val <= low || node.Val >= high {
		return false
	}
	// both the left and right subtree must also be valid
	return valid(node.Left, low, node.Val) && valid(node.Right, node.Val, high)
}

// given a number of nodes and a group of edges connecting those nodes determine if isDAG
// DFS stack
func validTree(n int, edges [][]int) bool {
	if len(edges)+1 != n {
		return false
	}
	// build an adjacency list
	list := make([][]int, n)
	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
		list[edge[1]] = append(list[edge[1]], edge[0])
	}
	// push edges onto the stack and pop to examine
	stack := make([]int, 0)
	// use a set to ignore nodes we have already seen
	seenset := make(map[int]bool)
	stack = push(stack, 0)
	seenset[0] = true

	for len(stack) > 0 {
		var node int
		stack, node = pop(stack)
		for _, v := range list[node] {
			if _, ok := seenset[v]; !ok {
				seenset[v] = true
				stack = push(stack, v)
			}
		}
	}

	return len(seenset) == n
}

// stack push
func push(stack []int, x int) []int {
	return append([]int{x}, stack...)
}

// stack pop
func pop(stack []int) ([]int, int) {
	t := stack[0]
	return stack[1:], t
}

// Recursive find depth
func maxDepth(root *TreeNode) int {
	depthLeft := 0
	depthRight := 0
	if root == nil {
		return 0
	}
	depthLeft = maxDepth(root.Left) + 1
	depthRight = maxDepth(root.Right) + 1
	return max(depthLeft, depthRight)
}

func max(l, r int) int {
	return int(math.Max(float64(l), float64(r)))
}

// BFS - Queue
func validPath(n int, edges [][]int, start int, end int) bool {
	// build an adjacency list
	list := make([][]int, n)
	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
		list[edge[1]] = append(list[edge[1]], edge[0])
	}
	nodes := make([]int, 0)
	seenset := make(map[int]bool)

	nodes = enqueue(nodes, start)
	seenset[start] = true

	for len(nodes) > 0 {
		var node int
		nodes, node = dequeue(nodes)
		if node == end {
			return true
		}
		for _, v := range list[node] {
			if _, ok := seenset[v]; !ok {
				seenset[v] = true
				nodes = enqueue(nodes, v)
			}
		}
	}
	return false
}

// add to back
func enqueue(nodes []int, n int) []int {
	return append([]int{n}, nodes...)
}

// remove from front
func dequeue(nodes []int) ([]int, int) {
	n := nodes[len(nodes)-1]
	return nodes[:len(nodes)-1], n
}
