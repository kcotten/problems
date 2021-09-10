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
