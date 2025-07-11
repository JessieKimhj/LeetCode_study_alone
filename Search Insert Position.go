package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1 //index of left, right end

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
