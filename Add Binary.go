package main

import "strconv"

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry
		if i >= 0 {
			digitA, _ := strconv.Atoi(string(a[i]))
			sum += digitA
			i--
		}
		if j >= 0 {
			digitB, _ := strconv.Atoi(string(b[j]))
			sum += digitB
			j--
		}

		result = strconv.Itoa(sum%2) + result // 현재 자리 추가
		carry = sum / 2                       // 올림수 계산
	}

	return result
}
