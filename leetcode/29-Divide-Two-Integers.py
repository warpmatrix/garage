def _div(dividend, divisor: int) -> int:
    if dividend < divisor:
        return 0
    cnt, minus = 1, divisor
    while minus + minus <= dividend:
        cnt += cnt
        minus += minus
    return cnt + _div(dividend - minus, divisor)

class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        intMax, intMin = 2 ** 31 - 1, -2 ** 31
        if dividend == 0:
            return 0
        if divisor == 1:
            return dividend
        if divisor == -1:
            if dividend == intMin:
                return intMax
            return -dividend
        sym = -(int((dividend < 0) ^ (divisor < 0)) * 2) + 1
        dividend, divisor = abs(dividend), abs(divisor)
        ret = sym * _div(dividend, divisor)
        return ret
