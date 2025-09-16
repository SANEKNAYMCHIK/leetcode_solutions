/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	middle, val := head, head
	for val != nil && val.Next != nil {
		middle = middle.Next
		val = val.Next.Next
	}

	prev, cur := (*ListNode)(nil), middle
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}