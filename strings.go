package main

func nextChar(ch byte, n int) byte {
	if ch++; ch > 'z' {
		return 'a'
	}
	return ch
}

func unique(s string) bool {
	memo := make(map[rune]int)
	for _, r := range s {
		if _, ok := memo[r]; !ok {
			memo[r]++
		} else {
			return false
		}
	}
	return true
}

func uniqueNoSpace(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func reverseCstyle(s string) string {
	if len(s) <= 2 {
		return s
	}
	b := make([]byte, len(s))
	size := len(s) - 2
	for i := 0; i < size; i++ {
		b[i], b[size-i] = s[size-i], s[i]
	}
	b[len(s)-1] = '\000'
	return string(b)
}

func isAnagram(s1, s2 string) bool {
	memo := make(map[rune]int)
	for _, r := range s1 {
		memo[r]++
	}
	for _, r := range s2 {
		if val, ok := memo[r]; ok {
			// consume the element
			if val >= 1 {
				memo[r]--
			} else {
				// we reached an element that has been exhausted
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// is s a subsequence of t, like a substring but not contiguous
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	j := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}
	return false
}
