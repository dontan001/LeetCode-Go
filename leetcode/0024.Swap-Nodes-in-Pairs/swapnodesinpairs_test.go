package leetcode

import "testing"

func TestSwapNodesInPairs(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
		head = SwapNodesInPairs(head)
		t.Logf("the swapped list: %+v", head)
	})
}
