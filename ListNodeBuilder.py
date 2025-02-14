from typing import List

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def buildListNode(items: List[int]) -> ListNode:
    ln = None
    items.reverse()  
    for i in items:
        ln = ListNode(i , ln)
    return ln