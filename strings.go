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
