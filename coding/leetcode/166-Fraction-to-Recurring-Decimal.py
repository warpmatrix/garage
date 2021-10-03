class Solution:
    def fractionToDecimal(self, numer: int, denom: int) -> str:
        ret = ""
        if (numer < 0) ^ (denom < 0):
            ret = "-"
        numer, denom = abs(numer), abs(denom)
        quotient, reminder = numer // denom , numer % denom
        if reminder == 0:
            if quotient == 0:
                return "0"
            return ret + str(quotient)
        ret += str(quotient) + "."

        rmdIdxs, index, frac = {}, 0, ""
        while reminder != 0 and reminder not in rmdIdxs:
            rmdIdxs[reminder], index, numer = index, index + 1, reminder * 10
            quotient, reminder = numer // denom, numer % denom
            frac += str(quotient)
        if reminder != 0:
            ret += frac[:rmdIdxs[reminder]] + "(" + frac[rmdIdxs[reminder]:] + ")"
        else:
            ret += frac
        return ret
