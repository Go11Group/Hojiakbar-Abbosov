def divide(dividend, divisor):
    INT_MIN = -2**31
    INT_MAX = 2**31 - 1

    if dividend == INT_MIN and divisor == -1:
        return INT_MAX
    
    negative = (dividend < 0) ^ (divisor < 0)
    
    dividend = abs(dividend)
    divisor = abs(divisor)
    
    quotient = 0
    while dividend >= divisor:
        temp_divisor = divisor
        multiple = 1
        while (temp_divisor << 1) <= dividend:
            temp_divisor <<= 1
            multiple <<= 1
        
        dividend -= temp_divisor
        quotient += multiple
    
    return -quotient if negative else quotient

dividend1, divisor1 = 10, 3
print("Input:", dividend1, divisor1)
print("Output:", divide(dividend1, divisor1))

dividend2, divisor2 = 7, -3
print("Input:", dividend2, divisor2)
print("Output:", divide(dividend2, divisor2))