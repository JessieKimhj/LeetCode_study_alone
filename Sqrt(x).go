package main

func mySqrt(x int) int {
	start, end := 0, x

	for start <= end { // 남은 탐색 범위를 다 찾아줌
		mid := (start + end) / 2

		if mid*mid > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start - 1
}
