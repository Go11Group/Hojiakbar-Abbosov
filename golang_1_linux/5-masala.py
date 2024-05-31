def removeElement(nums, val):
    if len(nums) == 0:
        return 0
    
    next_non_val = 0
    
    for i in range(len(nums)):
        if nums[i] != val:
            nums[next_non_val] = nums[i]
            next_non_val += 1
    
    return next_non_val

nums1 = [3, 2, 2, 3]
val1 = 3
print(removeElement(nums1, val1))  

nums2 = [0, 1, 2, 2, 3, 0, 4, 2]
val2 = 2
print(removeElement(nums2, val2))  
