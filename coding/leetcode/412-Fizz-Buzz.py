class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        ret: List[str] = ["" for _ in range(n)]
        for i in range(n):
            num = i + 1
            if num % 3 == 0:
                ret[i] += "Fizz"
            if num % 5 == 0:
                ret[i] += "Buzz"
            if len(ret[i]) == 0:
                ret[i] += str(num)
        return ret
