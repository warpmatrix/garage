class MedianFinder:
    def __init__(self):
        # max heap
        self.minHalf = list()
        # min heap
        self.maxHalf = list()

    def addNum(self, num: int) -> None:
        if len(self.maxHalf) == len(self.minHalf):
            heapq.heappush(self.maxHalf, -heapq.heappushpop(self.minHalf, -num))
        else:
            heapq.heappush(self.minHalf, -heapq.heappushpop(self.maxHalf, num))
        # if not self.maxHalf or self.maxHalf[0] < num:
        #     heapq.heappush(self.maxHalf, num)
        #     if len(self.maxHalf) > len(self.minHalf) + 1:
        #         heapq.heappush(self.minHalf,
        #             -heapq.heappop(self.maxHalf))
        # else:
        #     heapq.heappush(self.minHalf, -num)
        #     if len(self.minHalf) > len(self.maxHalf):
        #         heapq.heappush(self.maxHalf,
        #             -heapq.heappop(self.minHalf))

    def findMedian(self) -> float:
        return self.maxHalf[0] if len(self.maxHalf) > len(self.minHalf) \
            else (self.maxHalf[0] + -self.minHalf[0]) / 2



# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
