import sys
sys.path.append('../../../LeetCode')

from typing import Optional
from ListNodeBuilder import ListNode, buildListNode

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode], carry: int = 0) -> Optional[ListNode]:
        dm = divmod((l1.val if l1 else 0) + (l2.val if l2 else 0) + carry, 10)

        if l1 or l2:
            return ListNode(dm[1], self.addTwoNumbers(l1.next if l1 else None, l2.next if l2 else None, dm[0]))
        else:            
            return None if dm[1] == 0 else ListNode(dm[1], None)
        
sol = Solution()
#Test cases go here
print(f"{sol.addTwoNumbers(buildListNode([2,4,3]), buildListNode([5,6,4])) == buildListNode([7,0,8])}")
print(f"{sol.addTwoNumbers(buildListNode([0]), buildListNode([0])) == buildListNode([0])}")
print(f"{sol.addTwoNumbers(buildListNode([9,9,9,9,9,9,9]), buildListNode([9,9,9,9])) == buildListNode([8,9,9,9,0,0,0,1])}")