class Solution:
    def arrangeCoins(self, n: int) -> int:
        return int((pow(8 * n + 1, 0.5) - 1) / 2)
