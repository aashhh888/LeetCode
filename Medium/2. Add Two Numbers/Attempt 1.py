import sys
sys.path.append('../../../LeetCode')

from typing import Optional
from ListNodeBuilder import ListNode, buildListNode

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        res =  None
        
sol = Solution()
#Test cases go here
print(f"{sol.addTwoNumbers(buildListNode([2,4,3]), buildListNode([5,6,4])) == None if True else False}")