# def reverse(x):
#     sign = -1 if x < 0 else 1
#     x = abs(x)
#     result = 0
    
#     while x > 0:
#         digit = x % 10
#         result = result * 10 + digit
#         x //= 10
    
#     result *= sign
    
#     if result < -231 or result > 231 - 1:
#         return 0
    
#     return result

# print(reverse(123))  
# print(reverse(-123))  
# print(reverse(120))  

def reverse(x):
    INT_MAX = 2**31 - 1
    
    result = 0
    sign = 1 if x >= 0 else -1
    x = abs(x)
    
    while x > 0:
        if result > (INT_MAX - x % 10) // 10: 
            return 0
        result = result * 10 + x % 10
        x //= 10
    
    return result * sign

print(reverse(123))  
print(reverse(-123)) 
print(reverse(120))  