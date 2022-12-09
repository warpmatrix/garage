class Solution:
    def checkRecord(self, s: str) -> bool:
        a, l = 0, 0
        for _, ch in enumerate(s):
            if ch == 'A':
                a += 1
                if a >= 2:
                    return False
            if ch == 'L':
                l += 1
                if l >= 3:
                    return False
            else:
                l = 0
        return True
