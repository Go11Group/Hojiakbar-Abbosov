def letterCombinations(digits):
    if not digits:
        return []

    nums = {
        '2': 'abc',
        '3': 'def',
        '4': 'ghi',
        '5': 'jkl',
        '6': 'mno',
        '7': 'pqrs',
        '8': 'tuv',
        '9': 'wxyz'
    }

    def nokia(index, path):
        if index == len(digits):
            combinations.append(''.join(path))
            return

        for letter in nums[digits[index]]:
            path.append(letter)
            nokia(index + 1, path)
            path.pop()

    combinations = []
    nokia(0, [])
    return combinations

print(letterCombinations("23"))  
print(letterCombinations(""))   
print(letterCombinations("2"))  
