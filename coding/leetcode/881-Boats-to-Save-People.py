class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        lptr, rptr, ret = 0, len(people) - 1, 0
        while lptr <= rptr:
            if people[rptr] + people[lptr] <= limit:
                lptr, rptr = lptr + 1, rptr - 1
            else:
                rptr -= 1
            ret += 1
        return ret
