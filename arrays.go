package main

import (
	"reflect"
	"sort"
)

// Pascal's triangle
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := []int{}
		prevRow := result[i-1]
		row = append(row, 1)

		for j := 1; j < i; j++ {
			row = append(row, (prevRow[j-1] + prevRow[j]))
		}

		row = append(row, 1)
		result[i] = row
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	if target == 0 {
		if zeros(nums) {
			return [][]int{{0, 0, 0, 0}}
		}
	}
	results := make([][]int, 0)
	size := len(nums)
	sort.Ints(nums)
	for i := 0; i < size-3; i++ {
		for j := i + 1; j < size-3; j++ {
			left, right := i+2, size-1
			for left < right {
				temp := sum(nums[i], nums[j], nums[left], nums[right])
				if temp == target {
					result := []int{nums[i], nums[j], nums[left], nums[right]}
					if check(results, result) {
						results = append(results, result)
					}
					left = right
				} else if temp < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return results
}

func sum(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

// true to add array, false to skip
func check(results [][]int, result []int) bool {
	for _, quad := range results {
		if reflect.DeepEqual(quad, result) {
			return false
		}
	}

	return true
}

func zeros(nums []int) bool {
	for _, num := range nums {
		if num == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func kSum(nums []int, target, start, k int) [][]int {
	result := make([][]int, 0)
	if start == len(nums) || nums[start]*k > target || target > nums[len(nums)-1]*k {
		return result
	}
	if k == 2 {
		return twoSum(nums, target, start)
	}

	for i := start; i < len(nums); i++ {
		if i == start || nums[i-1] != nums[i] {
			for _, subset := range kSum(nums, target-nums[i], i+1, k-1) {
				subset = append([]int{nums[i]}, subset...)
				result = append(result, subset)
				//result[len(result)-1] = append(result[len(result)-1], subset...)
			}
		}

	}
	return result
}

// return index
func twoSum(nums []int, target, start int) [][]int {
	left, right := start, len(nums)-1
	result := make([][]int, 0)
	for left < right {
		currSum := nums[left] + nums[right]
		if currSum < target || (left > start && nums[left] == nums[left-1]) {
			left++
		} else if currSum > target || (right < len(nums)-1 && nums[right] == nums[right+1]) {
			right--
		} else {
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
		}
	}
	return result
}

// return index
func twoSumIdx(nums []int, target int) []int {
	idx1, idx2 := 0, 0
	memo := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := memo[target-nums[i]]; ok && i != val {
			idx1 = i
			idx2 = val
		} else {
			memo[nums[i]] = i
		}
	}
	return []int{idx1, idx2}
}

// backtrack
func combinationSum(candidates []int, target int) [][]int {
	results := make([][]int, 0)
	combo := make([]int, 0)
	var backtrack func(int, int, []int)
	backtrack = func(remain int, start int, combo []int) {
		if remain == 0 {
			// use temp buffer due to next loop operation modifying combo and therefore result
			temp := make([]int, len(combo))
			copy(temp, combo)
			results = append(results, temp)
			return
		} else if remain < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			combo = append(combo, candidates[i])
			// 1st call: target - the item we are currently looking at
			// Subsequently: subtract until we reach the target at beginning of func
			backtrack(remain-candidates[i], i, combo)
			combo = combo[:len(combo)-1]
		}
	}

	backtrack(target, 0, combo)
	return results
}

// Note: rewrite func to save zeros until end and consider them then
// otherwise flipped zeros will be considered to be a zero in the original matrix
// one approach is to make a copy of matrix and change the zeros there
// while looking at zeros in the original
func zeroOutMatrix(m [][]int) [][]int {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] == 0 {
				m[i] = zeroRow(m[i])
				m = zeroCol(m, j)
			}
		}
	}
	return m
}

func zeroRow(row []int) []int {
	for i := range row {
		row[i] = 0
	}
	return row
}

func zeroCol(m [][]int, j int) [][]int {
	for _, col := range m {
		col[j] = 0
	}
	return m
}

func rotate(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			t := matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[i][j]
			matrix[i][j] = t
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = t
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = t
		}
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	M, N := len(image), len(image[0])
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	var dfs func(int, int)
	dfs = func(r, c int) {
		if image[r][c] == color {
			image[r][c] = newColor
			if r >= 1 {
				dfs(r-1, c)
			}
			if r+1 < M {
				dfs(r+1, c)
			}
			if c >= 1 {
				dfs(r, c-1)
			}
			if c+1 < N {
				dfs(r, c+1)
			}
		}
	}
	dfs(sr, sc)
	return image
}

func maxAreaOfIsland(grid [][]int) int {
	M, N := len(grid), len(grid[0])
	seen := make([][]bool, M)
	result := 0
	for i := range grid {
		seen[i] = make([]bool, N)
	}

	var area func(int, int) int
	area = func(r, c int) int {
		if r < 0 || c < 0 || r >= M || c >= N || grid[r][c] == 0 || seen[r][c] {
			return 0
		}
		seen[r][c] = true
		return (1 + area(r+1, c) + area(r-1, c) + area(r, c-1) + area(r, c+1))
	}

	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			result = max(result, area(r, c))
		}
	}

	return result
}
