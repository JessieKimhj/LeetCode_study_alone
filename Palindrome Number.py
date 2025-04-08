class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        num_str = str(x)
    
    # Check if the string is equal to its reverse
        if num_str == num_str[::-1]:
            return True
    
        return False

        """
        letter = str(x)

        if x < 0 :
            return False
        elif letter == letter[::-1] :
            return letter == letter[::-1]
        return False
        """