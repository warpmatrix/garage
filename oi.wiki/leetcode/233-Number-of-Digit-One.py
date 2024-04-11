class Solution:
    def countDigitOne(self, n: int) -> int:
        ret, weight = 0, 1
        while weight <= n:
            high, cur, low = n // (weight * 10), n // weight % 10, n % weight
            ret += high * weight
            if cur > 1:
                ret += weight
            elif cur == 1:
                ret += low + 1
            weight *= 10
        return ret
