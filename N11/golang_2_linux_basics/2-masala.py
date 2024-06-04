def lengthOfLongestSubstring(s: str) -> int:
        if not s:
            return 0
        
        char_set = set()
        start = end = 0
        max_length = 0

        while end < len(s):
            if s[end] not in char_set:
                char_set.add(s[end])
                max_length = max(max_length,end - start + 1)
                end += 1
            else:
                char_set.remove(s[start])
                start += 1

        return max_length

print(lengthOfLongestSubstring("abcabcbb"))