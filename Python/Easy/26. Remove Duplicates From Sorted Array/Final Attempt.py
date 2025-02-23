from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        cur = 1
        for i in range(1, len(nums)):
            if nums[i] != nums[i-1]:
                nums[cur] = nums[i]
                cur = cur + 1
        return cur

sol = Solution()
#Test cases go here
case1 = [1,1,2]
print(f"{sol.removeDuplicates(case1) == 2}")
print(f"{case1 == [1,2,2]}")

case2 = [0,0,1,1,1,2,2,3,3,4]
print(f"{sol.removeDuplicates(case2) == 5}")
print(f"{case2 == [0,1,2,3,4,2,2,3,3,4]}")
