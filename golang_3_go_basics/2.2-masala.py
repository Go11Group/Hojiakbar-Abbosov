def generateParenthesis(n):
    def reverse(s='', left = 0, right = 0):
        if len(s) == 2 * n:
            result.append(s)
            return
        if left < n:
            reverse(s + '(', left + 1, right)
        if right < left:
            reverse(s + ')', left, right + 1)

    result = []
    reverse()
    return result


print(generateParenthesis(3))  
print(generateParenthesis(1))  
