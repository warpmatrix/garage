class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src, dst, k: int) -> int:
        dp = [0 if i == src else float("inf") for i in range(n)]
        for _ in range(k + 1):
            ndp = [cost for cost in dp]
            for f, t, p in flights:
                ndp[t] = min(ndp[t], dp[f] + p)
            dp = ndp
        return dp[dst] if dp[dst] != float("inf") else -1
