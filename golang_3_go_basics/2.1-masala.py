class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def removeNthFromEnd(head, n):
    length = 0
    curr = head
    while curr:
        length += 1
        curr = curr.next

    pos_to_remove = length - n

    if pos_to_remove == 0:
        return head.next

    prev = head
    for i in range(pos_to_remove - 1):
        prev = prev.next

    prev.next = prev.next.next

    return head


def print_linked_list(head):
    curr = head
    while curr:
        print(curr.val, end=" ")
        curr = curr.next
    print()

head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)
head.next.next.next.next = ListNode(5)

print("Original List:")
print_linked_list(head)

n = 2
head = removeNthFromEnd(head, n)

print("List after removing nth node from the end:")
print_linked_list(head)
