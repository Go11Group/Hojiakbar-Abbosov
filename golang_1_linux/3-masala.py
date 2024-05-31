def longest_common_prefix(strs):
    if not strs:  
        return ""

    min_len_str = min(strs, key=len)
    
    for i, char in enumerate(min_len_str):
        for other_str in strs:
            if other_str[i] != char:
                return min_len_str[:i]  
    
    return min_len_str  

print(longest_common_prefix(["flower","flow","flight"]))  
print(longest_common_prefix(["dog","racecar","car"]))  
