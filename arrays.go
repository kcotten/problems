package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
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
	fmt.Println(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[i][n-1-j]
			matrix[i][n-1-j] = t
		}
	}
	fmt.Println(matrix)
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

func updateMatrix(mat [][]int) [][]int {
	M, N := len(mat), len(mat[0])
	if M == 0 {
		return mat
	}
	// declare result
	dist := make([][]int, M)
	for i := range dist {
		dist[i] = make([]int, N)
	}
	// fill the result with max int - overflow?
	dist = fill(dist, MAX-100000)
	// left / up
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					dist[i][j] = min(dist[i][j], (dist[i-1][j] + 1))
				}
				if j > 0 {
					dist[i][j] = min(dist[i][j], (dist[i][j-1] + 1))
				}
			}
		}
	}
	// right / down
	for i := M - 1; i > -1; i-- {
		for j := N - 1; j > -1; j-- {
			if i < M-1 {
				dist[i][j] = min(dist[i][j], (dist[i+1][j] + 1))
			}
			if j < N-1 {
				dist[i][j] = min(dist[i][j], (dist[i][j+1] + 1))
			}
		}
	}

	return dist
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func fill(nums [][]int, a int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			nums[i][j] = a
		}
	}
	return nums
}

// merge intervals with 2d sort example
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	i := 0
	for i < len(intervals) {
		left, right := intervals[i][0], intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= right {
			right = max(right, intervals[j][1])
			j++
		}
		res = append(res, []int{left, right})
		i = j
	}

	return res
}

func exist(board [][]byte, word string) bool {
	ret := false
	size := len(word)
	M, N := len(board), len(board[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(int, int, int) bool
	dfs = func(i, j, length int) bool {
		if length == size {
			return true
		}
		if i < 0 || j < 0 || i >= M || j >= N || board[i][j] != word[length] {
			return false
		}
		// seed the board behind us to prevent following the same path
		temp := board[i][j]
		board[i][j] = '0'
		for _, dir := range directions {
			ret = ret || dfs(i+dir[0], j+dir[1], length+1)
		}
		board[i][j] = temp
		return ret
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

// reverse an integer
func reverse(x int) int {
	max := powInt(2, 31)
	fmt.Println(max)
	if x < 0 {
		z := x * -1
		ints := intToSlice(z)
		ret := sliceToInt(ints) * -1
		if ret < max*-1 {
			return 0
		} else {
			return ret
		}
	} else {
		ints := intToSlice(x)
		ret := sliceToInt(ints)
		if ret > (max - 1) {
			return 0
		} else {
			return ret
		}
	}
}

func intToSlice(n int) []int {
	ret := make([]int, 0)
	for n > 0 {
		ret = append(ret, n%10)
		n /= 10
	}
	return ret
}

func sliceToInt(s []int) int {
	ret := 0
	pow := 1
	for i := len(s) - 1; i >= 0; i-- {
		ret += s[i] * pow
		pow *= 10
	}
	return ret
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func reverseSlice(nums []int) []int {
	n := len(nums)
	for i := 0; i <= n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-i-1], nums[i]
	}
	return nums
}

var stacks []int

type stack struct {
	index int
	empty bool
}

func (s *stack) New(idx int) {
	s.index = idx
	s.empty = true
}

// three stacks using a single array
func threeStacks() {
	stacks = make([]int, 3)
	// one, two, three: keep track of the top of the stack
	// when stack 0 is updated increment 1 and 2
	// stack 1 push then increment 2
	stack0 := new(stack)
	stack1 := new(stack)
	stack2 := new(stack)
	stack0.New(0)
	stack1.New(1)
	stack2.New(2)
	fmt.Println(stack0, stack1, stack2)
	// push to 0
	if ok := stack0.threePush(1); ok {
		stack1.index++
		stack2.index++
	}
	if ok := stack0.threePush(3); ok {
		stack1.index++
		stack2.index++
	}
	if ok := stack1.threePush(4); ok {
		stack2.index++
	}
	stack2.threePush(7)
	fmt.Println(stacks)
}

func (s *stack) threePush(value int) bool {
	if s.empty {
		s.empty = !s.empty
		stacks[s.index] = value
		return false
	} else {
		stacks = append(stacks[:s.index+1], stacks[s.index:]...)
		stacks[s.index] = value
		return true
	}
}

func (s *stack) threePop() int {
	ret := stacks[s.index]
	stacks = append(stacks[:s.index], stacks[s.index+1:]...)
	return ret
}

// crypt[0] + crypt[1] = crypt[2] using mapping in solution
// no leading zeros
func isCryptSolution(crypt []string, solution [][]string) bool {
	sol := make(map[string]string, len(solution))
	res := make(map[string]string, len(solution))
	val1, val2 := "", ""
	for _, v := range solution {
		sol[v[0]] = v[1]
		res[v[1]] = v[0]
	}

	for _, v := range crypt[0] {
		s := string(v)
		val1 = val1 + sol[s]
	}
	for _, v := range crypt[1] {
		s := string(v)
		val2 = val2 + sol[s]
	}
	if !(checkLeading(val1) && checkLeading(val2)) {
		return false
	}
	fmt.Println(val1, val2)
	i, _ := strconv.Atoi(val1)
	j, _ := strconv.Atoi(val2)
	fmt.Println(i, j)
	// result as a string of ints
	result := strconv.Itoa(i + j)
	fmt.Println(result)
	if len(result) != len(crypt[2]) {
		return false
	}
	for i, r := range result {
		s := string(r)
		s2 := string(crypt[2][i])
		if sol[s2] != s {
			return false
		}
	}

	return true
}

func checkLeading(s string) bool {
	if len(s) == 1 {
		return true
	}
	t := strings.TrimLeft(s, "0")
	return t == s
}

// 0 <= index <= len(a)
func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
