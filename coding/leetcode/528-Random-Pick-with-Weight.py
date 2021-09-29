class Solution:
    def __init__(self, w: List[int]):
        self.preSum = list(accumulate(w))

    def pickIndex(self) -> int:
        weight = random.randint(1, self.preSum[len(self.preSum) - 1])
        return bisect.bisect_left(self.preSum, weight)

# Your Solution object will be instantiated and called as such:
# obj = Solution(w)
# param_1 = obj.pickIndex()
