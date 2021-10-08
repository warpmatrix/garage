class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        ans, length = [], 10
        cnts = defaultdict(int)
        for i in range(len(s) - length + 1):
            substr = s[i : i + length]
            cnts[substr] += 1
            if cnts[substr] == 2:
                ans.append(substr)
        return ans
