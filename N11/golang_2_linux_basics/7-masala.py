def maxArea(height):
    n = 0
    left, right = 0, len(height) - 1

    while left < right:
        h = min(height[left], height[right])
        n = max(n, h * (right - left))

        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return n


print(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))  
print(maxArea([1, 1]))  
