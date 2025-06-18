package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		ans := target - num
		if index, found := numMap[ans]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return nil
}
