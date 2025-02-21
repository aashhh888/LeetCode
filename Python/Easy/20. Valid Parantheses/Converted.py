from typing import List

class Solution:
    def isValid(self, s: str) -> bool:  
        if len(s) % 2 == 1:
            return False

        stack: List[str] = []
        
        for c in s:
            if (c == "(" or c == "[" or c == "{"):
                stack.append(c)
            elif (not len(stack) == 0) and ((c == ")" and stack[-1] == "(") 
                or (c == "]" and stack[-1] == "[") 
                or (c == "}" and stack[-1] == "{")):
                    stack.pop()
            else:
                return False        
        
        return len(stack) == 0

sol = Solution()
#Test cases go here
print(f"{sol.isValid("()") == True}")
print(f"{sol.isValid("()[]{}") == True}")
print(f"{sol.isValid("(]") == False}")
print(f"{sol.isValid("([])") == True}")
print(f"{sol.isValid("]") == False}")
print(f"{sol.isValid("){") == False}")
print(f"{sol.isValid("([}}])") == False}")
print(f"{sol.isValid("(){}}{") == False}")