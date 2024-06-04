def removeDuplicates(nums):
    if len(nums) == 0:
        return 0
    
    next_unique = 1
    
    for i in range(1, len(nums)):
        if nums[i] != nums[next_unique - 1]:
            nums[next_unique] = nums[i]
            next_unique += 1
    
    return next_unique

nums1 = [1, 1, 2]
print(removeDuplicates(nums1))  

nums2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
print(removeDuplicates(nums2))  
