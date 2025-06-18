package leetcode

import "testing"

func TestTwoAdd(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		one := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
		another := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.Next {
			t.Logf("%d", cur.Val)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		one := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
		another := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}}}}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.Next {
			t.Logf("%d", cur.Val)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		one := &ListNode{Val: 5, Next: nil}
		another := &ListNode{Val: 7, Next: nil}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.Next {
			t.Logf("%d", cur.Val)
		}
	})
}
