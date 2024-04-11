#include <vector>
using std::vector;

//  Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
   public:
    vector<ListNode *> splitListToParts(ListNode *head, int k) {
        int len = 0;
        for (ListNode *node = head; node != nullptr; node = node->next) {
            len++;
        }
        int quotient = len / k, reminder = len % k;
        std::vector<ListNode *> ret(k, nullptr);
        ListNode *node = head;
        for (int i = 0; i < ret.size(); i++) {
			ret[i] = node;
            int cnt = (i < reminder) ? quotient + 1 : quotient;
            for (int j = 0; j < cnt - 1; j++) {
                if (node == nullptr) break;
                node = node->next;
            }
			if (node == nullptr) break;
			ListNode *next = node->next;
			node->next = nullptr;
			node = next;
        }
		return ret;
    }
};
