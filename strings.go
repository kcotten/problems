package main

func nextChar(ch byte, n int) byte {
	if ch += 1; ch > 'z' {
		return 'a'
	}
	return ch
}
