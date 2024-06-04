class ListNode:
    def init(self, val=0, next=None):
        self.val = val
        self.next = next

def swapPairs(head):
    n = ListNode(0)
    n.next = head
    current = n

    while current.next and current.next.next:
        first = current.next
        second = current.next.next
        first.next = second.next
        current.next = second
        current.next.next = first
        current = current.next.next

    return n.next

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

print("Input:")
printList(head1)

result = swapPairs(head1)

print("Output:")
printList(result)