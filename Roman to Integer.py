class Solution(object):
    def romanToInt(self, s):
        Roman = {"I" : 1, "V" : 5, "X" : 10, "L" : 50, "C" : 100, "D" : 500, "M" : 1000}
    
        sum = 0

        for i in range(len(s) - 1):
            if Roman[s[i]] < Roman[s[i + 1]]:
                sum -= Roman[s[i]]
            else:
                sum += Roman[s[i]]

        sum += Roman[s[-1]]
        return sum
    