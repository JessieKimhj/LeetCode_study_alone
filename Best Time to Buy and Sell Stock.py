class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        두 리스트의 빼기값이 큰 순서대로 비교.
        brute force 알고리즘
        ascend = sorted.(prices)
        descend = sorted.(prices, reverse = True)
        for p in prices :
            for high_val in descend : 
                for low_val in ascend :
        """
        max_profit = 0
        min_price = float('inf')
        for price in prices:
            min_price = min(min_price, price)
            max_profit = max(max_profit, price - min_price)
        return max_profit