class Solution:
    def toHex(self, num: int) -> str:
        if num == 0:
            return "0"
        if num < 0:
            num += 2 ** 32
        digits = '0123456789abcdef'
        revRet = ""
        while num:
            revRet += digits[num & 0xf]
            num >>= 4
        return revRet[::-1]
