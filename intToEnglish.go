package main

var singles = []string{
	"Zero",
	"One",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
}

var twoNotTwenty = []string{
	"Ten",
	"Eleven",
	"Twelve",
	"Thirteen",
	"Fourteen",
	"Fifteen",
	"Sixteen",
	"Seventeen",
	"Eighteen",
	"Nineteen",
}

var tens = []string{
	"Iota",
	"Ten",
	"Twenty",
	"Thirty",
	"Forty",
	"Fifty",
	"Sixty",
	"Seventy",
	"Eighty",
	"Ninety",
}

func two(num int) string {
	switch {
	case num == 0:
		return ""
	case num < 10:
		return singles[num]
	case num < 20:
		return twoNotTwenty[num-10]
	default:
		tenner := num / 10
		rem := num - tenner*10
		if rem != 0 {
			return tens[tenner] + " " + singles[rem]
		} else {
			return tens[tenner]
		}
	}
}

func three(num int) string {
	hundred := num / 100
	rem := num - hundred*100
	result := ""
	switch {
	case hundred*rem != 0:
		result = singles[hundred] + " Hundred " + two(rem)
	case hundred == 0 && rem != 0:
		result = two(rem)
	case hundred != 0 && rem == 0:
		result = singles[hundred] + " Hundred"
	}
	return result
}

func numberToWords(num int) string {
	MAX := 2147483647
	if num < 0 || num > MAX {
		return ""
	}
	if num == 0 {
		return "Zero"
	}
	result := ""

	billion := num / 1000000000
	million := (num - billion*1000000000) / 1000000
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	rest := num - billion*1000000000 - million*1000000 - thousand*1000

	if billion != 0 {
		result += three(billion) + " Billion"
	}
	if million != 0 {
		if result != "" {
			result += " "
		}
		result += three(million) + " Million"
	}
	if thousand != 0 {
		if result != "" {
			result += " "
		}
		result += three(thousand) + " Thousand"
	}
	if rest != 0 {
		if result != "" {
			result += " "
		}
		result += three(rest)
	}

	return result
}

// func intToSlice(n int) []int {
// 	ret := make([]int, 0)
// 	for n > 0 {
// 		ret = append(ret, n%10)
// 		n /= 10
// 	}
// 	return ret
// }
