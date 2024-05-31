def isPalindrome(x: int) -> bool:
        n = str(x)
        return n[::-1] == n

print(isPalindrome(121))  
print(isPalindrome(-121))  
print(isPalindrome(10))