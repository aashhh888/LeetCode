from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        curr = 101
        for i in range(len(nums) - 1, -1, - 1):
            if curr == nums[i]:
                nums.pop(i)
            else:
                curr = nums[i]
        return len(nums)

sol = Solution()
#Test cases go here
case1 = [1,1,2]
print(f"{sol.removeDuplicates(case1) == 2}")
print(f"{case1 == [1,2]}")

case2 = [0,0,1,1,1,2,2,3,3,4]
print(f"{sol.removeDuplicates(case2) == 5}")
print(f"{case2 == [0,1,2,3,4]}")