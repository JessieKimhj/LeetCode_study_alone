class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        """
        계산한 숫자들은 str으로 변환해서 쪼갠다음 리스트로 내보내기
        """
        """
        l = len(digits)
        num = 0
        r = 0
        for i in range(l-1, -1, -1) :
            num += digits[i] * 10**r
            r +=1
        
        ans = num + 1
        return list(map(int, str(ans)))

        
        <2nd answer>
        """
        l = len(digits)

        for i in range(l-1, -1, -1):
            if digits[i] < 9 :
                digits[i] += 1
                return digits
            else : 
                digits[i] = 0
        return [1] + digits