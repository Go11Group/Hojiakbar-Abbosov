def roman_to_int(s):
    roman_dict = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    result = 0
    prev_value = 0

    for symbol in s[::-1]:  
        curr_value = roman_dict[symbol]
        if curr_value < prev_value:
            result -= curr_value
        else:
            result += curr_value
        prev_value = curr_value

    return result

print(roman_to_int("III"))  
print(roman_to_int("LVIII"))  
print(roman_to_int("MCMXCIV"))  
