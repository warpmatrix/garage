class Solution:
    def reverseVowels(self, s: str) -> str:
        s = list(s)
        lptr, rptr = 0, len(s) - 1
        vowels = "aeiouAEIOU"
        while lptr < rptr:
            while lptr < len(s) and s[lptr] not in vowels:
                lptr += 1
            while rptr >= 0 and s[rptr] not in vowels:
                rptr -= 1
            if lptr < rptr:
                s[lptr], s[rptr] = s[rptr], s[lptr]
                lptr, rptr  =  lptr + 1, rptr - 1
        return "".join(s)
