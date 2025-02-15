from typing import List

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
    def __eq__(self, other): 
        if not isinstance(other, ListNode):
            # Don't attempt to compare against unrelated types
            return NotImplemented

        areEqual = False        
        while self or other:
            areEqual = (self.val if self else -1) == (other.val if other else -1)
            self = self.next if self else None
            other = other.next if other else None
        return areEqual
    
    def __str__(self):            
        asString = ""
        while self:
            asString = asString + str(self.val)
            self = self.next
        return asString

def buildListNode(items: List[int]) -> ListNode:
    ln = None
    items.reverse()  
    for i in items:
        ln = ListNode(i , ln)
    return ln