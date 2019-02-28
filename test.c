#include <stdio.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode* newList(int*,int);

int main() {
    int a1[3] = {1,2,4};
    int a2[4] = {1,2,3,4};
    struct ListNode* l1 =  newList(a1,3);
    struct ListNode* l2 =  newList(a2,4);


    struct ListNode* l3 = mergeTwoLists(l1,l2);

    while(l3 != NULL) {
        printf("%d-",l3->val);
        l3 = l3->next;
    }

}

struct ListNode* newList(int *arr,int i) {
    struct ListNode *head;
    struct ListNode *m = head;
    struct ListNode *t;
    int j = 0;
    for (j=0;j<i;j++) {
        t.val = arr[j];
        t.next = NULL;
        m.next = t;
        t = NULL;
        m = m->next;
    }
    return head;
}


struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode *m = l1;
    struct ListNode *n = l2;
    struct ListNode *temp;
    struct ListNode *t = temp;

    while (m != NULL && n != NULL) {
        if(m->val <= n->val) {
            printf("%d-",m->val);
            temp->next = m;
            temp->next->next = NULL;
            m = m->next;
        } else {
            printf("-|%d|-",n->val);
            temp->next = n;
            temp->next->next = NULL;
            n = n->next;
        }
    }
    if (m->next != NULL) {
        temp->next = m;
    }
    if (n->next != NULL) {
        temp->next = n;
    }
    
    // while (m->next != NULL || n->next != NULL) {
    //     if(m) {
    //         temp->next = m;
    //     } else if(!n) {
    //         temp->next = n;
    //     }else if (m->val <= n->val) {
    //         temp->next = m;
    //         if (m->next != NULL) {
    //             m = m->next;
    //         }
    //     } else {
    //         temp->next = n;
    //         if (n->next != NULL) {
    //             n = n->next;
    //         }
    //     }
    //     temp = temp->next;
    // }
  
    return t->next;
}