#include <stack>

class MyQueue {
   public:
    /** Initialize your data structure here. */
    std::stack<int> in, out;
    MyQueue() {}

    /** Push element x to the back of queue. */
    void push(int x) { in.push(x); }

    /** Removes the element from in front of queue and returns that element. */
    int pop() {
        if (out.empty()) {
            while (!in.empty()) {
                out.push(in.top());
                in.pop();
            }
        }
        int ret = out.top();
        out.pop();
        return ret;
    }

    /** Get the front element. */
    int peek() {
        if (out.empty()) {
            while (!in.empty()) {
                out.push(in.top());
                in.pop();
            }
        }
        return out.top();
    }

    /** Returns whether the queue is empty. */
    bool empty() {
        return in.empty() && out.empty();
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */