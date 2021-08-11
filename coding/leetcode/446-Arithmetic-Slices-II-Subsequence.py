class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        ret = 0
        dp = [defaultdict(int) for _ in nums]
        for i, num in enumerate(nums):
            for j in range(i):
                d = num - nums[j]
                ret += dp[j][d]
                dp[i][d] += dp[j][d] + 1
        return ret
