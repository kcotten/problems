package main

import "fmt"

// find the minimum number of substrings you can create without having the same letters repeating in each substring
// to solve this find the longest unique substrings and add to array
// return the array
// does NOT work, broken
func minUniqueSubstrings(s string) []string {
	result := make([]string, 0)
	for i := 0; i < len(s); i++ {
		set := make(map[rune]bool)
		// seed the map and check if last char
		if i == len(s)-1 {
			result = append(result, s[i:])
			break
		}
		//set[rune(s[i])] = true
		for j := i; j < len(s); j++ {
			v := rune(s[j])
			if _, ok := set[v]; ok {
				// end search
				fmt.Println("Append string...", i, j)
				result = append(result, s[i:j])
				i = j - 1
			} else {
				if j == len(s)-1 {
					// we've come to the end of the string and need to append
					fmt.Println("Append end of string...", i, j)
					result = append(result, s[i:])
					// if s[j] == s[j-1] {
					// 	result = append(result, string(s[j]))
					// }
					i = j
					break
				}
				set[v] = true
			}
		}
	}
	return result
}

func helper(s string, result []string) (int, []string) {
	idx := 0
	for i := 0; i < len(s)-1; i++ {
		set := make(map[rune]bool)
		// seed the map
		set[rune(s[i])] = true
		for j := i + 1; j < len(s); j++ {
			v := rune(s[j])
			if _, ok := set[v]; ok {
				// end search
				fmt.Println("Append string...", i, j)
				result = append(result, s[i:j])
				i = j
			} else {
				if j == len(s)-1 {
					fmt.Println("Append end of string...", i, j)
					result = append(result, s[i:])
					idx = j + 1
					break
				}
				set[v] = true
			}
		}
	}
	return idx, result
}

// works
func minUniqueSubstrings2(s string) []string {
	result := make([]string, 0)
	count := [26]int{}
	sb := []byte(s)
	it := len(sb) - 1
	for len(sb) > 0 && it >= 0 {
		ch := sb[it]
		if count[ch-'a'] == 1 {
			// dup: right becomes left, append substring, reset count
			substr := string(sb[it+1:])
			result = append(result, substr)
			sb = sb[:it+1]
			it = len(sb) - 1
			count = [26]int{}
		} else {
			fmt.Println(string(sb[it:]))
			// count the char, decrement left
			count[ch-'a'] = 1
			it--
		}
	}
	if len(sb) > 0 {
		fmt.Println(string(sb))
		result = append(result, string(sb))
	}
	return result
}
