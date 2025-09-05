/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curVal := head
	var prevVal *ListNode = nil
	for curVal != nil {
		nextVal := curVal.Next
		curVal.Next = prevVal
		prevVal = curVal
		curVal = nextVal
	}
	return prevVal
}