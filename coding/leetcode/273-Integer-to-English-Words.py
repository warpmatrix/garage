num2Word = {
    1:"One", 2:"Two", 3:"Three", 4:"Four", 5:"Five",
    6:"Six", 7:"Seven", 8:"Eight", 9:"Nine", 10:"Ten",
    11:"Eleven", 12:"Twelve", 13:"Thirteen", 14:"Fourteen", 15:"Fifteen",
    16:"Sixteen", 17:"Seventeen", 18:"Eighteen", 19:"Nineteen",
    20:"Twenty", 30:"Thirty", 40:"Forty", 50:"Fifty",
    60:"Sixty", 70:"Seventy", 80:"Eighty", 90:"Ninety",
}
segNums = [1000000000, 1000000, 1000]
segStrs = ["Billion", "Million","Thousand"]

class Solution:
    def parseSeg(self, num: int) -> str:
        ret = ""
        if num >= 100:
            ret += num2Word[num // 100] + " Hundred"
            num %= 100

        for key in reversed(num2Word):
            if num >= key:
                if len(ret) > 0:
                    ret += " "
                num -= key
                ret += num2Word[key]

        return ret


    def numberToWords(self, num: int) -> str:
        if num == 0:
            return "Zero"
        ret = ""
        for i, segNum in enumerate(segNums):
            if num < segNum:
                continue
            if len(ret) > 0:
                ret += " "
            ret += self.parseSeg(num // segNum) + " " + segStrs[i]
            num %= segNum
        
        return ret
