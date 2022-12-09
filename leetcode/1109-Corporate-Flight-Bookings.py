class Solution:
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        ret = [0] * (n + 2)
        for f, l, s in bookings:
            ret[f], ret[l+1] = ret[f] + s, ret[l+1] - s
        for i in range(1, len(ret)):
            ret[i] += ret[i - 1]
        return ret[1:-1]
