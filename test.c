#include <stdio.h>

typedef node struct {
    int num;
    node *next;
}

int main() {
    node *result = {
        num = 0;
        node *next = null;
    };

}

node *add(m *node,n *node,node *result) {
    node *i = m;
    node *j = n;
    node *r = result;
    node *temp;

    if(i->next != null) {
        i = i->next;
    }
    if(j->next != null) {
        j = j->next;
    }
    if (i->next != null && j->next != null) {
        temp = add(i,j,r);
        if (temp->num > 9) {
            r->num = r->num + 1;
            temp->num = temp->num - 10;
            r->next = temp;
        }
    }
    int sum;
    sum = i->num + j->num;
    if (sum > 9) {
        temp->num = sum;
    }
    return temp;
}