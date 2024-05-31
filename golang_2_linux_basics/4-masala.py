def convert(s, numRows):
    if numRows == 1 or len(s) <= numRows:
        return s
    
    row = [''] * numRows
    row_idx = 0
    direction = 1  
    
    for char in s:
        row[row_idx] += char
        row_idx += direction
        
        if row_idx == 0 or row_idx == numRows - 1:
            direction *= -1
    
    return ''.join(row)

print(convert("PAYPALISHIRING", 3))  
print(convert("PAYPALISHIRING", 4))  
print(convert("A", 1))  