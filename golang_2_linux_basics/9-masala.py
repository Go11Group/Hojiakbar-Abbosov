def threeSumClosest(nums, target):
    nums.sort()  
    summa = float('inf')

    for i in range(len(nums) - 2):
        left, right = i + 1, len(nums) - 1

        while left < right:
            total = nums[i] + nums[left] + nums[right]
            diff = abs(total - target)

            if diff < abs(summa - target):
                summa = total

            if total == target:
                return summa
            elif total < target:
                left += 1
            else:
                right -= 1

    return summa

print(threeSumClosest([-1, 2, 1, -4], 1))  
print(threeSumClosest([0, 0, 0], 1))  
