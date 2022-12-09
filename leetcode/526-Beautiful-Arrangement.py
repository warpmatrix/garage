class Solution:
    def countArrangement(self, n: int) -> int:
        dp = [0] * (1 << n)
        dp[0] = 1
        for mask in range(1, 1 << n):
            num = bin(mask).count("1")
            for idx in range(n):
                if mask & (1 << idx) and (
                    num % (idx + 1) == 0 or (idx + 1) % num == 0):
                    dp[mask] += dp[mask ^ (1 << idx)]
        return dp[(1 << n) - 1]
