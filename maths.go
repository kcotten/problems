package main

var numerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var toRomans = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var nums = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func intToRoman(num int) string {
	ret := ""
	for num > 0 {
		v := highestDecimal(num)
		ret += toRomans[v]
		num -= v
	}
	return ret
}

func highestDecimal(n int) int {
	for _, v := range nums {
		if v <= n {
			return v
		}
	}
	return 1
}

func romanToInt(s string) int {
	ret := 0
	// add final element
	ret += numerals[string(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		t1, t2 := string(s[i]), string(s[i+1])
		// if preceding numeral less than -> subtract from the total
		if numerals[t1] < numerals[t2] {
			ret -= numerals[t1]
		} else {
			ret += numerals[t1]
		}
	}

	return ret
}

func powerOfTwo(n int) bool {
	return n&(n-1) == 0
}
