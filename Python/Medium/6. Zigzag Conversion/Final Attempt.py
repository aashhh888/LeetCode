class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        loc = 0
        res = [""] * numRows
        zig = True
        for c in s:
            if not res[loc] == "":
                res[loc] = res[loc] + c
            else:
                res[loc] = c

            if zig:
                loc = loc + 1
            else:
                loc = loc - 1

            if loc == 0:
                zig = True
            elif loc == (numRows - 1):
                zig = False

        return ''.join(res)

sol = Solution()
#Test cases go here
print(f"{sol.convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR"}")
print(f"{sol.convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI"}")
print(f"{sol.convert("A", 1) == "A"}")
print(f"{sol.convert("AB", 1) == "AB"}")