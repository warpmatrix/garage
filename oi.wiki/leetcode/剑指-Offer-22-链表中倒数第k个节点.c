// Definition for singly-linked list.
struct ListNode {
    int val;
    struct ListNode *next;
};

typedef struct ListNode ListNode;

ListNode* getKthFromEnd(ListNode* head, int k) {
    ListNode *ret = head;
    for (int i = 0; i < k; i++) {
        head = head->next;
    }
    while (head) {
        head = head->next;
        ret = ret->next;
    }
    return ret;
}

// ListNode* getKthFromEnd(ListNode* head, int k) {
//     int len = 0;
//     for (ListNode *node = head; node; node = node->next) {
//         len++;
//     }
//     ListNode *ret = head;
//     for (int i = 0; i < len - k; i++) {
//         ret = ret->next;
//     }
//     return ret;
// }
