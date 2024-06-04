class ListNode:
    def init(self, val=0, next=None):
        self.val = val
        self.next = next

def reverseKGroup(head, k):
    def reverseSublist(start, end):
        prev, current = None, start
        while current != end:
            next_node = current.next
            current.next = prev
            prev = current
            current = next_node
        return prev

    dummy = ListNode(0)
    dummy.next = head
    prev_group_end = dummy

    while True:
        group_start = prev_group_end.next
        group_end = group_start
        for _ in range(k - 1):
            if group_end is None:
                return dummy.next  
            group_end = group_end.next

        if group_end is None:
            return dummy.next  

        next_group_start = group_end.next
        reversed_group_start = reverseSublist(group_start, next_group_start)
        prev_group_end.next = reversed_group_start
        group_start.next = next_group_start

        prev_group_end = group_start

    return dummy.next

def printList(head):
    current = head
    while current:
        print(current.val, end=" ")
        current = current.next
    print()

head1 = ListNode(1)
head1.next = ListNode(2)
head1.next.next = ListNode(3)
head1.next.next.next = ListNode(4)
head1.next.next.next.next = ListNode(5)

print("Input:")
printList(head1)

k1 = 2
result1 = reverseKGroup(head1, k1)

print("Output (k = 2):")
printList(result1)

k2 = 3
result2 = reverseKGroup(head1, k2)

print("Output (k = 3):")
printList(result2)