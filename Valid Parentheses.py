class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        pair = {")":"(", "]":"[", "}": "{"}
        stack = []
        for i in s : 
            if i in pair :
                if stack : 
                    last_element = stack.pop()
                else :
                    last_element = '#'
                if pair[i] != last_element :
                    return False
            else : 
                stack.append(i)
        return len(stack)==0