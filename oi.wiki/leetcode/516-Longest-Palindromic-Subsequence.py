class Solution:
    def longestPalindromeSubseq(self, s: str) -> int:
        dp = [[1] * (len(s) - i) for i in range(len(s))]
        for offset in range(1, len(s)):
            for idx, ch in enumerate(s[:-offset]):
                if ch == s[idx + offset]:
                    if offset == 1:
                        dp[idx][offset] = 2
                    else:
                        dp[idx][offset] = 2 + dp[idx + 1][offset - 2]
                else:
                    dp[idx][offset] = max(dp[idx][offset-1], 
                        dp[idx+1][offset-1])
        return dp[0][len(s)-1]
