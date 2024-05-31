def addDigits(num):
    if num == 0:
        return 0 
    return 1 + ((num - 1 ) % 9 )

print(addDigits(38))
print(addDigits(0))