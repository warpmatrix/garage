class Solution:
    def compress(self, chars: List[str]) -> int:
        nlen, head = 0, 0
        for idx, s in enumerate(chars):
            if idx == len(chars) - 1 or s != chars[idx + 1]:
                chars[nlen] = s
                nlen += 1
                num = idx - head + 1
                if num > 1:
                    numStr = str(num)
                    for ch in numStr:
                        chars[nlen] = ch
                        nlen += 1
                head = idx + 1
        return nlen
