def titleToNumber(columnTitle):
    n = 0
    base = ord('A') - 1  
    
    for char in columnTitle:
        value = ord(char) - base  
        n = n * 26 + value  
    
    return n

print(titleToNumber("A"))  
print(titleToNumber("AB"))  
print(titleToNumber("ZY"))  