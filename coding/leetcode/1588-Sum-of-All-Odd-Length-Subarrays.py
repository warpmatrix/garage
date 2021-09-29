class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        ret = 0
        def odd(num: int) -> int:
            return (num + 1) // 2
        def even(num: int) -> int:
            return num // 2 + 1
        for i, num in enumerate(arr):
            left, right = i, len(arr) - i - 1
            ret += num * (odd(left) * odd(right) + even(left) * even(right))
        return ret
