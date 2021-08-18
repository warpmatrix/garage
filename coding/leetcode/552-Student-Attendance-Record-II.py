class Solution:
    def checkRecord(self, n: int) -> int:
        MOD = 10**9 + 7
        dp = [[[0] * (3) for _ in range(2)] for _ in range(n + 1)]
        dp[0][0][0] = 1
        for idx in range(1, n + 1):
            # a for rec[idx]
            for l in range(3):
                dp[idx][1][0] = (dp[idx][1][0] + dp[idx - 1][0][l]) % MOD
            # l for rec[idx]
            for a in range(2):
                for l in range(1, 3):
                    dp[idx][a][l] = (dp[idx][a][l] + dp[idx - 1][a][l - 1]) % MOD
            # p for rec[idx]
            for a in range(2):
                for l in range(3):
                    dp[idx][a][0] = (dp[idx][a][0] + dp[idx - 1][a][l]) % MOD
        ret = 0
        for a in range(2):
            for l in range(3):
                ret = (ret + dp[n][a][l]) % MOD
        return ret
