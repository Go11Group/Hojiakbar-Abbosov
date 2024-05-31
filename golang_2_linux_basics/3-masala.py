def longestPalindrome(s):
    n = len(s)
    if n <= 1:
        return s
    
    dp = [[False] * n for _ in range(n)]
    start = max_length = 0
    
    for length in range(1, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if length == 1:
                dp[i][j] = True
            elif length == 2:
                dp[i][j] = s[i] == s[j]
            else:
                dp[i][j] = s[i] == s[j] and dp[i+1][j-1]
            
            if dp[i][j] and length > max_length:
                start = i
                max_length = length
    
    return s[start:start + max_length]

print(longestPalindrome("babad"))  
print(longestPalindrome("cbbd"))  