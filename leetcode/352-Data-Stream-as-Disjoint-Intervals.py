from typing import List
from sortedcontainers import SortedDict

class SummaryRanges:
    def __init__(self):
        self.intervals = SortedDict()
    

    def addNum(self, val: int) -> None:
        length = len(self.intervals)
        idx1 = self.intervals.bisect_right(val)
        idx0 = idx1 - 1 if idx1 > 0 else length
        lefts, rights = self.intervals.keys(), self.intervals.values()
        if idx0 != length and lefts[idx0] <= val <= rights[idx0]:
            return
        
        leftMerge = (idx0 != length) and rights[idx0] + 1 == val
        rightMerge = (idx1 != length) and lefts[idx1] - 1 == val
        if leftMerge and rightMerge:
            left, right = lefts[idx0], rights[idx1]
            self.intervals.popitem(idx1)
            self.intervals.popitem(idx0)
            self.intervals[left] = right
        elif leftMerge:
            self.intervals[lefts[idx0]] += 1
        elif rightMerge:
            right = rights[idx1]
            self.intervals.popitem(idx1)
            self.intervals[val] = right
        else:
            self.intervals[val] = val


    def getIntervals(self) -> List[List[int]]:
        return list(self.intervals.items())



# Your SummaryRanges object will be instantiated and called as such:
# obj = SummaryRanges()
# obj.addNum(val)
# param_2 = obj.getIntervals()
