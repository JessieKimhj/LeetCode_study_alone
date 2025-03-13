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

/*
type Script Solution
function mySqrt(x: number): number {
  // as per the constrain x lies from 0 - 2^31-1
  let low = 0;

  let high = x;

  while (low <= high) {
    const mid = Math.floor((low + high) / 2);

    const square = mid * mid;
    else high = mid - 1;
  }

  // at the end of the iteration our high
  // pointer will point to the correct number
  return high;
}

*/
