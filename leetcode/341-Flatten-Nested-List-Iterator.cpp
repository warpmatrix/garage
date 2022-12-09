#include <queue>
#include <vector>
using std::vector;

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
class NestedInteger {
   public:
    // Return true if this NestedInteger holds a single integer, rather than a
    // nested list.
    bool isInteger() const;

    // Return the single integer that this NestedInteger holds, if it holds a
    // single integer The result is undefined if this NestedInteger holds a
    // nested list
    int getInteger() const;

    // Return the nested list that this NestedInteger holds, if it holds a
    // nested list The result is undefined if this NestedInteger holds a single
    // integer
    const vector<NestedInteger> &getList() const;
};

class NestedIterator {
    std::queue<int> intList;
    void flattenList(const vector<NestedInteger> &nestedList) {
        for (auto &&nestedInt : nestedList) {
            if (nestedInt.isInteger()) {
                intList.push(nestedInt.getInteger());
            } else {
                flattenList(nestedInt.getList());
            }
        }
    }

   public:
    NestedIterator(vector<NestedInteger> &nestedList) {
        flattenList(nestedList);
    }

    int next() {
        int ret = intList.front();
        intList.pop();
        return ret;
    }

    bool hasNext() { return !intList.empty(); }
};

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i(nestedList);
 * while (i.hasNext()) cout << i.next();
 */