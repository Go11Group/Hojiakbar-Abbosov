def repeatedCharacter(s: str) -> str:
    dct = {}

    for i , j in enumerate(s):
        if j in dct:
            return j
        dct[j] = i

    return None

print(repeatedCharacter("abccbaacz"))
print(repeatedCharacter("abcdd"))