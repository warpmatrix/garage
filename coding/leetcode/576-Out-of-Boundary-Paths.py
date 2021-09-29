class Solution:
    def findPaths(self, m, n, maxMove, startRow, startColumn: int) -> int:
        MOD, ret = 10**9 + 7, 0

        dp = [[0] * n for _ in range(m)]
        dp[startRow][startColumn] = 1
        for _ in range(maxMove):
            dpNew = [[0] * n for _ in range(m)]
            for r in range(m):
                for c in range(n):
                    if do[r][c] == 0:
                        continue
                    for nr, nc in [(r - 1, c), (r + 1, c), (r, c - 1), (r, c + 1)]:
                        if 0 <= nr < m and 0 <= nc < n:
                            dpNew[nr][nc] = (dpNew[nr][nc] + dp[r][c]) % MOD
                        else:
                            ret = (ret + dp[r][c]) % MOD
            dp = dpNew
        
        return ret
