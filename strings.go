package main

func nextChar(ch byte, n int) byte {
	if ch++; ch > 'z' {
		return 'a'
	}
	return ch
}
