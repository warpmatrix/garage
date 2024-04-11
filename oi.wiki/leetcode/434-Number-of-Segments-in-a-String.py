class Solution:
    def countSegments(self, s: str) -> int:
        return len([str for str in s.split(' ') if len(str) > 0])
