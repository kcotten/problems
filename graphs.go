package main

import "fmt"

type pair struct {
	i, j int
}

func Len(q []pair) int {
	if len(q) == 0 {
		return 0
	}
	l := 0
	for i := 0; i < len(q); i++ {
		l = i
	}
	return l + 1
}

func orangesRotting(grid [][]int) int {
	q := make([]pair, 0)

	freshOranges := 0
	row, col := len(grid), len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				p := pair{i, j}
				q = enqueue(q, p)
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}

	q = enqueue(q, pair{-1, -1})
	elapsed := -1
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for Len(q) > 0 {
		var p pair
		//fmt.Printf("Q:%v, L:%v\n",q, Len(q))
		q, p = dequeue(q)
		r := p.i
		c := p.j

		if r == -1 {
			elapsed++
			if Len(q) > 0 {
				q = enqueue(q, pair{-1, -1})
			}
		} else {
			for _, dir := range directions {
				neighborRow := r + dir[0]
				neighborCol := c + dir[1]
				if neighborRow >= 0 && neighborRow < row &&
					neighborCol >= 0 && neighborCol < col {
					if grid[neighborRow][neighborCol] == 1 {
						grid[neighborRow][neighborCol] = 2
						freshOranges--
						q = enqueue(q, pair{neighborRow, neighborCol})
					}
				}
			}
		}
	}
	if freshOranges == 0 {
		return elapsed
	}
	return -1
}

func enqueue(q []pair, v pair) []pair {
	q = append(q, v)
	return q
}

func dequeue(q []pair) ([]pair, pair) {
	v := q[0]
	return q[1:], v
}

func canFinish(num int, prereq [][]int) bool {
	// memoize the courses with their dependencies (adjacencies)
	memo := make(map[int][]int)
	for _, v := range prereq {
		// revisit
		if val, ok := memo[v[1]]; ok {
			memo[v[1]] = append(val, v[0])
		} else {
			memo[v[1]] = append(val, v[0])
		}
	}
	fmt.Println(memo)
	path := make([]bool, num)
	for current := 0; current < num; current++ {
		if isCyclic(current, memo, path) {
			return false
		}
	}

	return true
}

func isCyclic(current int, memo map[int][]int, path []bool) bool {
	if path[current] {
		return true
	}
	if _, ok := memo[current]; !ok {
		return false
	}
	path[current] = true

	ret := false
	for _, v := range memo[current] {
		if ret = isCyclic(v, memo, path); ret {
			break
		}
	}
	path[current] = false
	return ret
}
