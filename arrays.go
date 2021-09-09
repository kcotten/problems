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
