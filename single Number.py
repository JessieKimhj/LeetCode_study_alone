class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        
        result = 0
        for num in nums:
            result ^= num
            print("result", result)
            print("num", num)
            print("--------------------")
        return result