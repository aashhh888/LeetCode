from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(0, len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:                    
                    return [i,j]

sol = Solution()
#Test cases go here
print(f"{sol.twoSum([2,7,11,15], 9) == [0,1]}")
print(f"{sol.twoSum([3,2,4], 6) == [1,2]}")
print(f"{sol.twoSum([3,3], 6) == [0,1]}")