// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode *reverseBetween(ListNode *head, int left, int right) {
    // node before front
    ListNode *begin = nullptr;
    for (int i = 1; i < left; i++) {
        if (i == 1)
            begin = head;
        else
            begin = begin->next;
    }
    ListNode *front = begin ? begin->next : head;
    ListNode *prev = front;
    ListNode *node = prev->next;
    for (int i = left + 1; i <= right; i++) {
        ListNode *next = node->next;
        node->next = prev;
        prev = node;
        node = next;
    }
    // after loop, node == end, prev == back
    front->next = node;
    if (begin) {
        begin->next = prev;
    } else {
        head = prev;
    }
    return head;
}
